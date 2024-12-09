package waitgroup

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func ApplyWaitGroup() {
	go first() // makes go routine here creates a thread
	wg.Add(1)  // add one more thread to the system lets wg know there is an extra thread runnning
	fmt.Println("GOARCH\t\t", runtime.GOARCH)
	fmt.Println("current cpu:\t", runtime.NumCPU())
	fmt.Println("current goroutines:", runtime.NumGoroutine())
	second()
	wg.Wait() // asks to wait until the goroutines are finished Done() method lets Wait know its finished
}

func first() {
	for i := 0; i < 10; i++ {
		fmt.Println("first waiting index: ", i)
	}
	wg.Done() // it is the await command, sends the okay when finished
}

func second() {
	for i := 0; i < 10; i++ {
		fmt.Println("second waiting index: ", i)
	}
}
