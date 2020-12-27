package code_blocking

import (
	"strconv"
)

type adapter struct{}

type adapterInterface interface {
	ReceiveData()
}

func NewAdapterTest() adapterInterface {
	return &adapter{}
}

func (a *adapter) ReceiveData() {
	dataChannel := make(chan []byte)
	receiver := []string{"AMQP", "DATABASE"}

	handler := NewHandler(receiver, dataChannel)
	go handler.ListenToAdapter()

	go func() {
		for i := 0; i <= 10; i++ {
			b := []byte(strconv.Itoa(i))
			dataChannel <- b
		}

		close(dataChannel)
	}()

	handler.Wg.Wait()
}
