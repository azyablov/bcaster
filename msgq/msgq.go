package msgq

import (
	"fmt"
	"sync"
)

type MsgQueue struct {
	msgArr []string
	size   int
	next   int
	last   int
	l      sync.Mutex
}

func NewMsgQueue(size int) *MsgQueue {
	m := MsgQueue{
		msgArr: make([]string, size),
		size:   size,
		next:   0,
		last:   0,
		l:      sync.Mutex{},
	}

	return &m
}

func (mq *MsgQueue) Enqueue(msg string) error {
	mq.l.Lock()
	if (mq.next+1)%mq.size == mq.last {
		return fmt.Errorf("")
	}
	mq.msgArr[mq.next] = msg
	mq.next = (mq.next + 1) % mq.size
	mq.l.Unlock()
	return nil
}

func (mq *MsgQueue) Dequeue() (string, error) {
	mq.l.Lock()
	if mq.last == mq.next {
		return "", fmt.Errorf("message queue is empty")
	}
	msg := mq.msgArr[mq.last]
	mq.last = (mq.last + 1) % mq.size
	mq.l.Unlock()
	return msg, nil
}
