package main

import (
	"fmt"
	"sync"
)

func InitWorkerPool(size int) int {
	var x int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			x++
			mu.Unlock()
			fmt.Println("x:", x)
		}()
	}

	wg.Wait()
	return x
}

func main() {
	fmt.Println("result:", InitWorkerPool(1000))
}
