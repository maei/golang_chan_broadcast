package code_blocking

import (
	"fmt"
	"log"
	"sync"
	"time"
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

/// DATABASE
type Database struct {
}

func (db *Database) Init(wg *sync.WaitGroup) chan []byte {
	recCh := make(chan []byte)
	go db.subscribe(recCh, wg)
	return recCh
}

func (db *Database) subscribe(ch chan []byte, wg *sync.WaitGroup) {
	for {
		data, open := <-ch
		if open {
			db.storeData(data)
		} else {
			wg.Done()
			return
		}
	}
}

func (db *Database) storeData(payload []byte) {
	log.Println(fmt.Sprintf("hello from Database:%v", payload))
}

/// AQMP
type AMQP struct {
}

func (amqp *AMQP) Init(wg *sync.WaitGroup) chan []byte {
	recCh := make(chan []byte)
	go amqp.subscribe(recCh, wg)
	return recCh
}

func (amqp *AMQP) subscribe(ch chan []byte, wg *sync.WaitGroup) {
	for {
		data, open := <-ch
		if open {
			amqp.publishAMQP(data)
		} else {
			wg.Done()
			return
		}
	}
}

func (amqp *AMQP) publishAMQP(payload []byte) {
	time.Sleep(time.Second * 2)
	log.Println(fmt.Sprintf("hello from AMQP:%v", payload))
}
