package main

import (
	"github.com/maei/golang_chan_broadcast/src/code_non_blocking_ch_buff"
	_ "net/http/pprof"
	"time"
)

func main() {

	// code_blocking
	//a := code_blocking.NewAdapterTest()
	//
	//a.ReceiveData()

	/*
		code_non_blocking
		b := code_non_blocking.NewAdapter()
		b.ReceiveData()
		time.Sleep(time.Second * 5)

		c := code_non_blocking.NewAdapter()
		c.ReceiveData()
	*/

	// code_non_blocking_ch_buff
	d := code_non_blocking_ch_buff.NewAdapter()
	d.ReceiveData()

	time.Sleep(time.Minute * 5)

}
