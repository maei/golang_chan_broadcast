package code_non_blocking

import (
	"container/list"
	"sync"
	"time"
)

type Buffer interface {
	Add(value []byte)
	Get() []byte
	Clear()
	Pop()
	IsEmpty() bool
	ConsumeBuffer(f func([]byte), timeout time.Duration)
}

type buffer struct {
	List  *list.List
	Mutex sync.Mutex
}

func NewBuffer() Buffer {
	return &buffer{
		List:  list.New(),
		Mutex: sync.Mutex{},
	}
}

/*callback function to observe changes in buffer
Consume data in FIFO and deletes consumed entry after it was computed*/
func (b *buffer) ConsumeBuffer(fn func([]byte), timeout time.Duration) {
	for {
		switch b.IsEmpty() {
		case false:
			fn(b.Get())
			b.Pop()
		case true:
			time.Sleep(timeout)
		}
	}
}

// Puts a new Value to the End of the Linked List
func (b *buffer) Add(value []byte) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	_ = b.List.PushBack(value)
}

// Deletes the first Value in the Linked List
func (b *buffer) Pop() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.List.Remove(b.List.Front())
}

// Get the first Value of the Linked List
func (b *buffer) Get() []byte {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.List.Front().Value.([]byte)
}

// Delete the whole Linked List
func (b *buffer) Clear() {
	b.List.Init()
}

// Check if Linked List is Empty
func (b *buffer) IsEmpty() bool {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.List.Len() == 0
}
