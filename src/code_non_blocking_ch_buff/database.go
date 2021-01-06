package code_non_blocking_ch_buff

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

func (db *Database) StoreData(dataUpstream []byte) {
	log.Println(fmt.Sprintf("data written to file:%v", dataUpstream))
	time.Sleep(time.Second * 1)
}
