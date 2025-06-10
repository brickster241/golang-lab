package learnlab

import (
	"fmt"
	"time"
)

func channelsIntro() {

	greet := make(chan string)
	greetString := "Sashiburidana, Mugiwara :))"
	fmt.Printf("Program Started at : %v\n", time.Now())
	
	// Needs to be placed inside a goroutine, blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
	go func() {
		time.Sleep(2 * time.Second)
		greet <- greetString
	}()

	receiver := <-greet
	fmt.Printf("%v -> Receiver : %s", time.Now(), receiver)

}

func buffChannelsBlockingOnReceive() {
	// ======== BLOCKING ON RECEIVE ONLYIF THE BUFFER IS EMPTY
	ch := make(chan int, 2)
	go func() {
		time.Sleep(time.Second)
		ch <- 1
		ch <- 2
	}()
	fmt.Println("Value1 :", <- ch)
	fmt.Println("Value2 :", <- ch)
	fmt.Println("End of Program :)")

}

func buffChannelsBlockingOnSend() {
	// ========== BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from Buffer ...")
	go func() {
		fmt.Println("Goroutine 2 second timer started ...")
		time.Sleep(3 * time.Second)
		fmt.Println("Received :", <-ch)
	}()

	fmt.Println("Blocking Starts....")
	ch <- 3 // Blocks because the buffer is full
	fmt.Println("Blocking ends....")
	fmt.Println("Received :", <- ch)
	fmt.Println("Received :", <- ch)
}