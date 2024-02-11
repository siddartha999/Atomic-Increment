package main

import (
	"fmt"
	"sync"
)
var mutex sync.Mutex
var wg sync.WaitGroup

var count int = 0

func incrementCount() {
	//Pessimistic locking
	mutex.Lock()
	count++
	mutex.Unlock()
	wg.Done()
}

func doCount() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go incrementCount()
	}
}

func main() {
	doCount()
	wg.Wait()
	fmt.Println("Atomic: Count value after 10K increment calls: %s", count)
}