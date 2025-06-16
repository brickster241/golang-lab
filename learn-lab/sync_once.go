package learnlab

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("This will not be repeated no matter how many times you call this :))")
}

func syncONCE() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("Goroutine #", i)
			once.Do(initialize)
		}()
	}
	wg.Wait()
}