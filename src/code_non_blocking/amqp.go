package code_non_blocking

import (
	"log"
	"sync"
	"time"
)

/// AQMP
type AMQP struct {
	Buffer Buffer
}

func (amqp *AMQP) Init(wg *sync.WaitGroup) chan []byte {
	recCh := make(chan []byte)
	amqp.Buffer = NewBuffer()
	go amqp.Buffer.ConsumeBuffer(amqp.publishData, time.Second*1)
	go amqp.subscribe(recCh, wg)
	return recCh
}

func (amqp *AMQP) subscribe(ch chan []byte, wg *sync.WaitGroup) {
	for {
		select {
		case data, open := <-ch:
			if open {
				log.Printf("amqp data received: %v", data)
				amqp.Buffer.Add(data)
			} else {
				wg.Done()
				log.Println("amqp data streaming done")
				return
			}
		}
	}
}

func (amqp *AMQP) publishData(payload []byte) {
	log.Printf("publish data over amqp: %v", payload)
	time.Sleep(time.Second * 2)

}
