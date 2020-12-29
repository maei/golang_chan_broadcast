package code_non_blocking

import (
	"container/list"
	"sync"
	"time"
)

type ByteBuffer interface {
	Add(value []byte)
	Get() []byte
	Clear()
	Pop()
	IsEmpty() bool
	ConsumeBuffer(f func([]byte), timeout time.Duration)
}

type byteBuffer struct {
	List  *list.List
	Mutex sync.Mutex
}

// Implements a double linked list
func NewBuffer() ByteBuffer {
	return &byteBuffer{
		List:  list.New(),
		Mutex: sync.Mutex{},
	}
}

// callback function to observe changes in byteBuffer
// Consumes data in FIFO and deletes consumed entry after it was computed*/
func (b *byteBuffer) ConsumeBuffer(fn func([]byte), timeout time.Duration) {
	for {
		switch b.IsEmpty() {
		case true:
			time.Sleep(timeout)
		case false:
			// sync function call
			fn(b.Get())
			// deletes element which was consumed by the receiver function
			b.Pop()
			//fmt.Printf("buffer size: %v\n", b.List.Len())
		}
	}
}

// Puts a new Value to the End of the Linked List
func (b *byteBuffer) Add(value []byte) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	_ = b.List.PushBack(value)
}

// Deletes the first Value in the Linked List
func (b *byteBuffer) Pop() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.List.Remove(b.List.Front())
}

// Get the first Value of the Linked List
func (b *byteBuffer) Get() []byte {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.List.Front().Value.([]byte)
}

// Delete the whole Linked List
func (b *byteBuffer) Clear() {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.List.Init()
}

// Check if Linked List is Empty
func (b *byteBuffer) IsEmpty() bool {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.List.Len() == 0
}
