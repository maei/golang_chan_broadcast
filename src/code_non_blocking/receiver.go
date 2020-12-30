package code_non_blocking

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type ReceiverInterface interface {
	InitReceiverNode(wg *sync.WaitGroup) chan []byte
	subscribe(ch chan []byte, wg *sync.WaitGroup)
}

type ReceiverNode struct {
	ProcessID             uuid.UUID
	ReceiverNodeName      string
	Buffer                ByteBuffer
	Callback              func([]byte)
	ReceiverNodeInterface interface{}
	Timeout               time.Duration
	BufferSize            int16
}

func ReceiverNodeFactory(receiverNodeName string) (*ReceiverNode, error) {
	if receiverNodeName == "AMQP" {
		rn := &ReceiverNode{
			ProcessID:             uuid.New(),
			ReceiverNodeInterface: NewAmqp(),
			ReceiverNodeName:      receiverNodeName,
			Timeout:               time.Millisecond * 500,
			BufferSize:            10,
		}
		rn.Callback = rn.ReceiverNodeInterface.(*AMQP).PublishData
		return rn, nil
	}
	if receiverNodeName == "DATABASE" {
		rn := &ReceiverNode{
			ProcessID:             uuid.New(),
			ReceiverNodeInterface: NewDatabase(),
			ReceiverNodeName:      receiverNodeName,
			Timeout:               time.Millisecond * 500,
			BufferSize:            10,
		}
		rn.Callback = rn.ReceiverNodeInterface.(*Database).StoreData
		return rn, nil
	}
	return &ReceiverNode{}, errors.New(fmt.Sprintf("receiver: %v not known", receiverNodeName))
}

func (rec *ReceiverNode) InitReceiverNode(wg *sync.WaitGroup) chan []byte {
	recCh := make(chan []byte, rec.BufferSize)
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
