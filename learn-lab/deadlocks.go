package learnlab

import (
	"fmt"
	"sync"
	"time"
)

func deadlocks() {
	var mu1, mu2 sync.Mutex
	go func(){
		mu1.Lock()
		fmt.Println("Goroutine 1 Locked Mu1")
		time.Sleep(time.Second)

		mu2.Lock()
		fmt.Println("Goroutine 1 Locked Mu2")

		mu1.Unlock()
		mu2.Unlock()
	}()
	go func(){
		mu2.Lock()
		fmt.Println("Goroutine 2 Locked Mu2")
		time.Sleep(time.Second)

		mu1.Lock()
		fmt.Println("Goroutine 2 Locked Mu1")

		mu2.Unlock()
		mu1.Unlock()
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main Function completed..")
}