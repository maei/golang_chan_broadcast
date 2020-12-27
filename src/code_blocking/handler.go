package code_blocking

import "sync"

type Handler struct {
	Receiver     []ReceiverInterface
	ReceiverChan []chan []byte
	Data         chan []byte
	Wg           sync.WaitGroup
}

type HandlerInterface interface {
	ListenToAdapter()
}

func NewHandler(newReceiver []string, data chan []byte) *Handler {
	h := &Handler{
		Data: data,
	}
	for _, v := range newReceiver {
		h.Receiver = append(h.Receiver, ReceiverFactory(v))
	}

	h.Wg.Add(len(h.Receiver) + 1)

	for _, t := range h.Receiver {
		h.ReceiverChan = append(h.ReceiverChan, t.Init(&h.Wg))
	}

	return h
}

func (hi *Handler) ListenToAdapter() {
	for {
		data, open := <-hi.Data
		if open {
			for _, x := range hi.ReceiverChan {
				x <- data
			}
		} else {
			for _, z := range hi.ReceiverChan {
				close(z)
			}
			hi.Wg.Done()
			return
		}
	}
}
