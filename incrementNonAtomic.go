package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var count int = 0

func incrementCount() {
	count++
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
	doCount()
	wg.Wait()
	fmt.Println("NonAtomic: Count value after 10K increment calls: %s", count)
}

