package code_non_blocking_ch_buff

import (
	"log"
	"sync"
)

type UpstreamHandlerInterface interface {
	ListenToUpstream()
}

type UpstreamHandler struct {
	ReceiverNodes []*ReceiverNode
	DataUpstream  chan []byte
	Wg            sync.WaitGroup
}

func NewUpstreamHandler(receiverNodes []string, bufferSize int8) *UpstreamHandler {
	handler := &UpstreamHandler{
		DataUpstream: make(chan []byte, bufferSize),
	}

	for _, val := range receiverNodes {
		node, nErr := ReceiverNodeFactory(val)
		if nErr != nil {
			log.Println(nErr)
		} else {
			handler.ReceiverNodes = append(handler.ReceiverNodes, node)
		}
	}

	handler.Wg.Add(len(handler.ReceiverNodes) + 1)

	for _, val := range handler.ReceiverNodes {
		val.InitReceiverNode()
	}
	return handler
}

func (h *UpstreamHandler) ListenToUpstream() {
	for {
		select {
		case data, open := <-h.DataUpstream:
			if open {
				for _, x := range h.ReceiverNodes {
					log.Printf("%v data received: %v, with ID: %v", x.ReceiverNodeName, data, x.ProcessID.String())
					x.Buffer.Add(data)
				}
			} else {
				h.Wg.Done()
				return
			}
		}
	}
}
