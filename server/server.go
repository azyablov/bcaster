package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/azyablov/bcaster/msgq"
	bc "github.com/azyablov/bcaster/protos/bcaster"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	host = flag.String("host", "localhost", "IP@ or FQDN used to spawn listener, by default localhost.")
	port = flag.Int("port", 8088, "Listening port, by default will use 8088.")
)

type bServer struct {
	bc.UnimplementedBCasterServer
	msgQ        *msgq.MsgQueue
	bpm         BackpressureMeter
	cList       map[string]int32
	cListLock   sync.Mutex
	cListRegCh  chan string
	cTimeoutMax int32
}

func (s *bServer) RegisterClient(ctx context.Context, r *bc.RegistrationRequest) (*bc.RegistrationAccept, error) {
	client := r.GetName()
	s.cListRegCh <- client
	return &bc.RegistrationAccept{Name: client, Timeout: s.cTimeoutMax}, nil
}

func (s *bServer) Registrar() {
	for cName := range s.cListRegCh {

		s.cListLock.Lock()
		if _, ok := s.cList[cName]; !ok {
			// Regiser new client
			log.Printf("New client registration request from: %q with timeout %v", cName, s.cTimeoutMax)
			s.cList[cName] = s.cTimeoutMax
		} else {
			// Refeshing registration
			s.cList[cName] = s.cTimeoutMax
			log.Printf("Refereshing registration for: %q", cName)
		}
		s.cListLock.Unlock()
	}
}

func (s *bServer) timeUpdater() {
	for {
		time.Sleep(1 * time.Second)
		s.cListLock.Lock()
		for k := range s.cList {
			s.cList[k] -= 1
			// Removing expired registrations
			if s.cList[k] <= 0 {
				log.Printf("Registration has been expired for client: %s", k)
				delete(s.cList, k)
			}
		}
		s.cListLock.Unlock()
	}
}

func (s *bServer) checkRegistration(ctx context.Context) error {
	// Playing with context
	var cName string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// Getting client name from MD
		mdSl := md.Get("name")
		if len(mdSl) == 0 {
			return fmt.Errorf("regisration has been expired or no client found")
		}
		cName = mdSl[0]
	}

	// Checking registration
	s.cListLock.Lock()
	defer s.cListLock.Unlock()
	if _, ok := s.cList[cName]; !ok {
		return fmt.Errorf("regisration has been expired or no client found")
	}
	return nil
}

func (s *bServer) SendShortMsg(ctx context.Context, msg *bc.ShortMsg) (*bc.DeliveryStatus, error) {

	// Playing with context
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// MD handling
		for key, value := range md {
			log.Printf("%s: %q", key, value)
		}
	}

	ds := bc.DeliveryStatus{Enq: false}
	select {
	case <-s.bpm.bp:
		defer func() {
			s.bpm.bp <- struct{}{}
		}()
		m := msg.GetMsg()
		// Logging
		log.Printf("Received message: %v", m)
		log.Printf("Number of tockens: %v", len(s.bpm.bp))

		// Application business logic
		m, err := businessLogic(m)
		if err != nil {
			return &ds, err
		}
		// ==========================

		// Processing
		if err := s.msgQ.Enqueue(m); err != nil {
			return &ds, err
		}
		ds.Enq = true
		return &ds, nil
	case <-ctx.Done():
		return &ds, ctx.Err()
	default:
		return &ds, BackpressureError{s: "not enougth capacity"}
	}
}

func businessLogic(msg string) (string, error) {
	// Our nice business with this service
	switch msg {
	case "Gravity":
		msg = fmt.Sprintf("%s Falls", msg)
		return msg, nil
	case "Bill":
		msg = fmt.Sprintf("%s Cipher", msg)
		return msg, nil
	default:
		return "", fmt.Errorf("bad request to the service")
	}
}

func (s *bServer) RecvShortMsg(ctx context.Context, dr *bc.DeliveryRequest) (*bc.ShortMsg, error) {
	// Uncomment to demonstate how context works on client side
	// time.Sleep(10 * time.Second)
	select {
	case <-s.bpm.bp:
		// Getting tocken back
		defer func() {
			s.bpm.bp <- struct{}{}
		}()

		// Checking registration
		if err := s.checkRegistration(ctx); err != nil {
			return nil, err
		}

		// Getting message from the service queue
		msg, err := s.msgQ.Dequeue()
		if err != nil {
			return nil, err
		}

		// Creting log record
		log.Printf("Sending message: %v", msg)

		return &bc.ShortMsg{Msg: msg}, nil

	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return nil, BackpressureError{s: "not enougth capacity"}
	}
}

// func (s *bServer) getMsgForClient(ctx context.Context) (string, error) {
// 	return "", nil
// }

func newBServer(nRegReqMax int, simReqMax int, cRegTimeoutMax int32, msgQsize int) *bServer {
	var bSrv bServer = bServer{
		msgQ:        msgq.NewMsgQueue(msgQsize),
		bpm:         *NewBackpressureMeter(simReqMax),
		cList:       make(map[string]int32),
		cListLock:   sync.Mutex{},
		cListRegCh:  make(chan string, nRegReqMax),
		cTimeoutMax: cRegTimeoutMax,
	}
	log.Print("Starting timeUpdater()...")
	go bSrv.timeUpdater()
	go bSrv.Registrar()
	return &bSrv
}

func main() {
	// Parsing CLI params
	flag.Parse()
	// Creating TCP socket
	sock, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create new gRPC server w/o concrete service registered
	srv := grpc.NewServer()

	// Service instance object
	svc := newBServer(10, 10, 10, 10)

	bc.RegisterBCasterServer(srv, svc)

	log.Printf("server listening at %v", sock.Addr())
	if err := srv.Serve(sock); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Backpressure

type BackpressureError struct {
	s string
}

func (e BackpressureError) Error() string {
	return fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), e.s)
}

type BackpressureMeter struct {
	bp chan struct{}
}

func NewBackpressureMeter(limit int) *BackpressureMeter {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &BackpressureMeter{bp: ch}
}
