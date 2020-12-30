package main

import (
	"github.com/maei/golang_chan_broadcast/src/code_non_blocking"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	//a := code_blocking.NewAdapterTest()
	//
	//a.ReceiveData()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	b := code_non_blocking.NewAdapter()
	b.ReceiveData()
	time.Sleep(time.Second * 5)

	c := code_non_blocking.NewAdapter()
	c.ReceiveData()

	time.Sleep(time.Minute * 5)

}
