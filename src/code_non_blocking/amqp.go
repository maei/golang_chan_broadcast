package code_non_blocking

import (
	"log"
	"time"
)

/// AQMP
type AMQP struct {
}

type AmqpInterface interface {
	PublishData(payload []byte)
}

func NewAmqp() AmqpInterface {
	return &AMQP{}
}

func (amqp *AMQP) PublishData(payload []byte) {
	log.Printf("publish data over amqp: %v", payload)
	time.Sleep(time.Second * 1)

}
