package channelblock

import (
	"fmt"
	"sync"
)

func ExecFanout() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sendFanOut(even, odd)
	go receiveFanOut(even, odd, fanin)
	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func receiveFanOut(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}

func sendFanOut(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}
