package code_non_blocking

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/// DATABASE
type Database struct {
	Buffer Buffer
}

func (db *Database) Init(wg *sync.WaitGroup) chan []byte {
	recCh := make(chan []byte)
	db.Buffer = NewBuffer()
	go db.Buffer.ConsumeBuffer(db.storeData, 1)
	go db.subscribe(recCh, wg)
	return recCh
}

func (db *Database) subscribe(ch chan []byte, wg *sync.WaitGroup) {
	for {
		select {
		case data, open := <-ch:
			if open {
				log.Printf("database data received: %v", data)
				db.Buffer.Add(data)
			} else {
				wg.Done()
				log.Println("database data streaming done")
				return
			}
		}
	}
}

func (db *Database) storeData(payload []byte) {
	log.Println(fmt.Sprintf("data written to file:%v", payload))
	time.Sleep(time.Second * 2)
}
