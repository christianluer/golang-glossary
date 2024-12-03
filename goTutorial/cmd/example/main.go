package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Time taken: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)
}

func dbCall(i int) {
	// simulate a db call
	var delay float32 = 2000
	time.Sleep(1 * time.Duration(delay) * time.Millisecond)
	fmt.Println("the result from the db is: ", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}
