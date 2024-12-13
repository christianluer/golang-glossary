package channelblock

import "fmt"

func ExecSelectCommaOk() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	//send
	go sendSelectCommaOk(eve, odd, quit)
	//receive
	receiveSelectCommaOk(eve, odd, quit)
	close(eve)
	close(odd)
	fmt.Println("about to exit")
}

func sendSelectCommaOk(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receiveSelectCommaOk(e, o <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd  channel:", v)
		case i, ok := <-quit: // checks if a channel is closed
			if !ok {
				fmt.Println("from comma ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}

func SimpleExecCommaOk() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
}

func SimpleExecCommaOkClosed() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}
