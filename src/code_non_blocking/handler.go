package code_non_blocking

import (
	"log"
	"sync"
)

type UpstreamHandlerInterface interface {
	ListenToUpstream()
}

type UpstreamHandler struct {
	ReceiverNodes     []*ReceiverNode
	ReceiverNodesChan []chan []byte
	DataUpstream      chan []byte
	Wg                sync.WaitGroup
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
		handler.ReceiverNodesChan = append(handler.ReceiverNodesChan, val.InitReceiverNode(&handler.Wg))
	}

	return handler
}

func (h *UpstreamHandler) ListenToUpstream() {
	for {
		select {
		case data, open := <-h.DataUpstream:
			if open {
				for _, x := range h.ReceiverNodesChan {
					x <- data
				}
			} else {
				for _, z := range h.ReceiverNodesChan {
					close(z)
				}
				h.Wg.Done()
				return
			}
		}
	}
}
