package code_non_blocking

import (
	"fmt"
	"log"
	"time"
)

/// DATABASE
type Database struct {
}

type DatabaseInterface interface {
	StoreData(payload []byte)
}

func NewDatabase() DatabaseInterface {
	return &Database{}
}

func (db *Database) StoreData(payload []byte) {
	log.Println(fmt.Sprintf("data written to file:%v", payload))
	time.Sleep(time.Second * 4)
}
