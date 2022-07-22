package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	bc "github.com/azyablov/bcaster/protos/bcaster"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	host = flag.String("host", "localhost", "IP@ or FQDN used to connect to target, by default localhost.")
	port = flag.Int("port", 8088, "Target port, by default will use 8088.")
	name = flag.String("name", "Deeper", "name of our client")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%v", *host, *port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't not connect to the host %v: %v", *host, err)
	}
	defer conn.Close()

	client := bc.NewBCasterClient(conn)

	// Working with context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Altering metadata w/ name of the client (potentially could be UUID/tocken)
	ctx = metadata.AppendToOutgoingContext(ctx, "name", *name)

	// Registering ourslef with service as a client
	_, err = client.RegisterClient(ctx, &bc.RegistrationRequest{Name: *name})
	if err != nil {
		log.Fatalf("can't register: %v", err)
	}

	// Sending short msg via service
	delStatus, err := client.SendShortMsg(ctx, &bc.ShortMsg{Msg: "Gravity"})
	if err != nil {
		log.Fatalf("could not send a message: %v", err)
	}
	// Maybe it's not necessary in fact, but to demonstate gRPC functionality
	if delStatus.Enq {
		log.Println("Message put in the queue")
	}
	// Trying to send incorrect request to demonstate a business logic
	//delStatus, err = client.SendShortMsg(ctx, &bc.ShortMsg{Msg: "Robbie"})
	delStatus, err = client.SendShortMsg(ctx, &bc.ShortMsg{Msg: "Bill"})

	if err != nil {
		// Handling it gracefully that that
		log.Printf("could not send a message: %v", err)
	} else {
		if delStatus.Enq {
			log.Println("Message put in the queue")
		}
	}

	// Spawning || go procedure to receive and print message
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		msg, err := client.RecvShortMsg(ctx, &bc.DeliveryRequest{})
		if err != nil {
			log.Fatalf("could not receive a message: %v", err)
		}

		log.Printf("Receiving message back: %s", msg.Msg)
	}()

	go func() {
		defer wg.Done()
		// The same as previous one, but with delay
		time.Sleep(5 * time.Second)

		msg, err := client.RecvShortMsg(ctx, &bc.DeliveryRequest{})
		if err != nil {
			log.Fatalf("could not receive a message: %v", err)
		}

		log.Printf("Receiving message back: %s", msg.Msg)
	}()

	// Giving 15 seconds to receive messages back and cancelling context

	time.Sleep(15 * time.Second)
	cancel()
	wg.Wait()

}
