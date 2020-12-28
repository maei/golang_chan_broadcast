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
	h := &Handler{
		DataUpstream: upstreamData,
	}

	for _, v := range newReceiverModes {
		h.ReceiverNodes = append(h.ReceiverNodes, ReceiverNodeFactory(v))
	}

	h.Wg.Add(len(h.ReceiverNodes) + 1)

	for _, t := range h.ReceiverNodes {
		h.ReceiverNodesChan = append(h.ReceiverNodesChan, t.InitReceiverNode(&h.Wg))
	}

	return h
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
