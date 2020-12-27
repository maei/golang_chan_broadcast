package code_non_blocking

import (
	"sync"
)

type ReceiverInterface interface {
	Init(wg *sync.WaitGroup) chan []byte
}

func ReceiverFactory(receiverName string) ReceiverInterface {
	if receiverName == "DATABASE" {
		return &Database{}
	}
	if receiverName == "AMQP" {
		return &AMQP{}
	}
	return nil
}
