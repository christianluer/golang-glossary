package channelblock

import "fmt"

// <- blocks until it receives or sends something
// so ExecuteWrongWay is being blocked in line 12 because it does not reach anywhere
// If we use a goroutine, the Anonymous functions gets blocked and the main thread receives it so it both can go on
// in line 21 <-c it blocks the main code until it receives the c<-42 signal and then keeps running, same with the goroutine in sending the info

func ExecuteWrongWay() {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}

func Execute() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

func ExecuteWithBufferChannel() {
	c := make(chan int, 2) // you stablish a buffer for 2 interactions the number of interactions can be >= of the value of real interactions and could work
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func ExecuteWithBufferChannelWrongWay() {
	c := make(chan int, 1) // only one buffer, but two blocks are called
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
