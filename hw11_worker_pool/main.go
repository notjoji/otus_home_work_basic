package main

import (
	"fmt"
	"sync"
)

func InitWorkerPool(size int, isPrinted bool) int {
	var x int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			x++
			if isPrinted {
				fmt.Println("x:", x)
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	return x
}

func main() {
	fmt.Println("result:", InitWorkerPool(1000, true))
}
