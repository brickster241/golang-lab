package learnlab

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signals() {
	pid := os.Getpid()
	fmt.Println("Process Id :", pid)
	sigs := make(chan os.Signal, 1)

	// Notify channel on interrupt or terminate signal
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		sig := <- sigs
		fmt.Println("Received Signal :", sig)
		fmt.Println("Graceful exit ....")
		os.Exit(0)
	}()

	// Simulate some work
	fmt.Println("Working....")
	for {
		time.Sleep(time.Second)
	}
}