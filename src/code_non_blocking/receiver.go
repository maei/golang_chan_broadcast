package code_non_blocking

import (
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type ReceiverInterface interface {
	InitReceiverNode(wg *sync.WaitGroup) chan []byte
}

/// ReceiverNode
type ReceiverNode struct {
	ProcessID             uuid.UUID
	ReceiverNodeName      string
	Buffer                Buffer
	Callback              func([]byte)
	ReceiverNodeInterface interface{}
	Timeout               time.Duration
}

func ReceiverNodeFactory(receiverName string) *ReceiverNode {
	if receiverName == "AMQP" {
		rn := &ReceiverNode{
			ProcessID:             uuid.New(),
			ReceiverNodeInterface: NewAmqp(),
			ReceiverNodeName:      "AMQP",
			Timeout:               time.Millisecond * 500,
		}
		rn.Callback = rn.ReceiverNodeInterface.(*AMQP).PublishData
		return rn
	}
	if receiverName == "DATABASE" {
		rn := &ReceiverNode{
			ProcessID:             uuid.New(),
			ReceiverNodeInterface: NewDatabase(),
			ReceiverNodeName:      "DATABASE",
			Timeout:               time.Millisecond * 500,
		}
		rn.Callback = rn.ReceiverNodeInterface.(*Database).StoreData
		return rn
	}
	return &ReceiverNode{}
}

func (rec *ReceiverNode) InitReceiverNode(wg *sync.WaitGroup) chan []byte {
	recCh := make(chan []byte)
	rec.Buffer = NewBuffer()
	go rec.Buffer.ConsumeBuffer(rec.Callback, rec.Timeout)
	go rec.subscribe(recCh, wg)
	return recCh
}

func (rec *ReceiverNode) subscribe(ch chan []byte, wg *sync.WaitGroup) {
	for {
		select {
		case data, open := <-ch:
			if open {
				log.Printf("%v data received: %v, with ID: %v", rec.ReceiverNodeName, data, rec.ProcessID.String())
				rec.Buffer.Add(data)
			} else {
				wg.Done()
				log.Printf("%v data streaming done for ID: %v", rec.ReceiverNodeName, rec.ProcessID.String())
				return
			}
		}
	}
}
