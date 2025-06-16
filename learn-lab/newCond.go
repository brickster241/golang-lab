package learnlab

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 7

type buffer struct {
	items []int
	mu    sync.Mutex
	cond *sync.Cond
}

// Creates a new buffer
func newBuffer(size int) *buffer {
	bf := &buffer{
		items: make([]int, 0, size),
	}

	bf.cond = sync.NewCond(&bf.mu)
	return bf
}

func (b* buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == bufferSize {
		// WE NEED TO WAIT FOR SOME ITEMS TO BE CONSUMED, AS IT IS FULL
		fmt.Println("Waiting for consumer to consume item.....")
		b.cond.Wait()	// Temporarily unlocks the mutex.
	}

	b.items = append(b.items, item)
	fmt.Println("Produced :", item)
	b.cond.Signal()	// Wakes up the other goroutine.
}

func (b* buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		// WAIT FOR PRODUCE FUNCTION TO DO SOMETHING.
		fmt.Println("Waiting for producer to produce item.....")
		b.cond.Wait() // Temporarily unlocks the mutex.
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed :", item)
	b.cond.Signal()		// Wakes up the other goroutine.
	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 100 {
		b.produce(i * 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 100 {
		b.consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func newCond() {

	bf := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(bf, &wg)
	go consumer(bf, &wg)

	wg.Wait()
}