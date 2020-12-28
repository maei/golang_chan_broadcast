package code_non_blocking

import (
	"sync"
)

type HandlerInterface interface {
	ListenToAdapter()
}

type Handler struct {
	ReceiverNodes     []*ReceiverNode
	ReceiverNodesChan []chan []byte
	DataUpstream      chan []byte
	Wg                sync.WaitGroup
}

func NewHandler(newReceiverModes []string, upstreamData chan []byte) *Handler {
	handler := &Handler{
		DataUpstream: upstreamData,
	}

	for _, val := range newReceiverModes {
		handler.ReceiverNodes = append(handler.ReceiverNodes, ReceiverNodeFactory(val))
	}

	handler.Wg.Add(len(handler.ReceiverNodes) + 1)

	for _, val := range handler.ReceiverNodes {
		handler.ReceiverNodesChan = append(handler.ReceiverNodesChan, val.InitReceiverNode(&handler.Wg))
	}

	return handler
}

func (h *Handler) ListenToAdapter() {
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
