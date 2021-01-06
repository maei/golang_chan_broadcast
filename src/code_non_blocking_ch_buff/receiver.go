package code_non_blocking_ch_buff

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type ReceiverInterface interface {
	InitReceiverNode(wg *sync.WaitGroup)
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

func (rec *ReceiverNode) InitReceiverNode() {
	log.Printf("%v buffer started", rec.ReceiverNodeName)
	rec.Buffer = NewBuffer()
	go rec.Buffer.ConsumeBuffer(rec.Callback, rec.Timeout)
}
