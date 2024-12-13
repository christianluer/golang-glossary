package channelblock

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func WgSimpleExecDirectionalChannelSender() {
	c := make(chan<- int, 2)
	c <- 42
	c <- 43
	fmt.Printf("%T\n", c)
}

func WgSimpleExecDirectionalChannelReceiver() {
	c := make(<-chan int, 2)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}

func WgExecMain() {
	c := make(chan int)
	wg.Add(2)
	go wgsend(c)
	go wgreceive(c)
	wg.Wait()
	fmt.Println("about to exit")
}

func wgsend(c chan<- int) {
	c <- 42
	wg.Done()
}
func wgreceive(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
