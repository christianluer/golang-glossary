package main

import (
	"fmt"
	waitgroup "go_tutorials/cmd/concurrency/waitGroup"
)

func main() {
	fmt.Println("hello")
	waitgroup.ApplyWaitGroup()
}
