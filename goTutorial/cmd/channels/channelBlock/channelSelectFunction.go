package channelblock

import "fmt"

func ExecSelect() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	//send
	go sendSelect(eve, odd, quit)
	//receive
	receiveSelect(eve, odd, quit)
	close(eve)
	close(odd)
	close(quit)
	fmt.Println("about to exit")
}

func sendSelect(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receiveSelect(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd  channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}
