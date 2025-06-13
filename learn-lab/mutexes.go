package learnlab

import (
	"fmt"
	"sync"
)

type counter struct {
	mu sync.Mutex
	value int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func mutexes() {
	var wg sync.WaitGroup
	ctr := &counter{}
	num_goroutines := 10

	fmt.Println("Initial Value :", ctr.value)
	// Final Value should be 10000, this is because of mutex
	for range num_goroutines {
		wg.Add(1)
		go func(){
			defer wg.Done()
			for range 1000 {
				ctr.increment()
			}
		}()

	}
	wg.Wait()
	fmt.Println("Program Finished..., Final Value :", ctr.value)
}