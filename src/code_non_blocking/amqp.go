package code_non_blocking

import (
	"log"
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

func (amqp *AMQP) PublishData(dataUpstream []byte) {
	log.Printf("publish data over amqp: %v", dataUpstream)
	//time.Sleep(time.Second * 1)

}
