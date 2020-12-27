package code_non_blocking

import (
	"fmt"
	"strconv"
)

type adapter struct{}

type adapterInterface interface {
	ReceiveData()
}

func NewAdapter() adapterInterface {
	return &adapter{}
}

// Receivers always block until there is data to receive.
// If the channel is unbuffered, the sender blocks until the receiver has received the value.
func (a *adapter) ReceiveData() {
	dataChannel := make(chan []byte)
	receiver := []string{"AMQP", "DATABASE"}

	handler := NewHandler(receiver, dataChannel)
	go handler.ListenToAdapter()

	go func() {
		for i := 0; i <= 10; i++ {
			b := []byte(strconv.Itoa(i))
			dataChannel <- b
			fmt.Println(b)
		}
		close(dataChannel)
	}()

	handler.Wg.Wait()
}
