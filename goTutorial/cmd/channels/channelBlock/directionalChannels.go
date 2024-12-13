package channelblock

import (
	"fmt"
)

func SimpleExecDirectionalChannelSender() {
	c := make(chan<- int, 2)
	c <- 42
	c <- 43
	fmt.Printf("%T\n", c)
}

func SimpleExecDirectionalChannelReceiver() {
	c := make(<-chan int, 2)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}

func ExecMain() {
	c := make(chan int)
	go send(c)
	receive(c) // goroutine in send will execute entirely because the receive is blocking the main thread until it receives something
	fmt.Println("about to exit")
}

func send(c chan<- int) {
	c <- 42
}
func receive(c <-chan int) {
	fmt.Println(<-c)
}
