package main

import (
	"fmt"
	"sync"
)

func InitWorkerPool(size int, print bool) int {
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
			if print {
				fmt.Println("x:", x)
			}
		}()
	}

	wg.Wait()
	return x
}

func main() {
	fmt.Println("result:", InitWorkerPool(1000, true))
}
