package channelblock

import "fmt"

func ExecRange() {
	c := make(chan int)
	go sendRange(c)
	for v := range c { // if the channel didnt close in sendRange then this will block the code waiting for another sendRange value
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func sendRange(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // if this is not closed, then the channel will keep open
}
