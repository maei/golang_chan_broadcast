package code_non_blocking

import (
	"fmt"
	"strconv"
)

type adapter struct{}

type AdapterInterface interface {
	ReceiveData()
}

func NewAdapter() AdapterInterface {
	return &adapter{}
}

// Receivers always block until there is data to receive.
// If the channel is unbuffered, the sender blocks until the receiver has received the value.
func (a *adapter) ReceiveData() {
	receiver := []string{"AMQP", "DATABASE", "asdasd"}

	handler := NewUpstreamHandler(receiver, 0)
	go handler.ListenToUpstream()

	go func() {
		for i := 0; i <= 10; i++ {
			b := []byte(strconv.Itoa(i))
			handler.DataUpstream <- b
			fmt.Println(b)
		}
		close(handler.DataUpstream)
	}()

	handler.Wg.Wait()
}
