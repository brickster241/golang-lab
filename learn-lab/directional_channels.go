package learnlab

import "fmt"

// PRODUCER - CONSUMER PATTERN
func main() {
	ch := make(chan int)

	produceData(ch)
	consumeData(ch)

}

// WE CAN ONLY SEND DATA INTO A CHANNEL
func produceData(ch chan <- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// WE CAN ONLY RECEIVE DATA FROM THE CHANNEL
func consumeData(ch <- chan int) {
	for value := range ch {
		fmt.Println("Received :", value)
	}
}

