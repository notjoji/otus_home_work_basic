package main

import (
	"fmt"
	"time"
)

func SensorGeneratorChannel() chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 60; i++ {
			time.Sleep(1 * time.Second)
			c <- i
		}
		close(c)
	}()
	return c
}

func AverageSumChannel(sensor <-chan int) chan float32 {
	out := make(chan float32)
	go func() {
		var sum float32
		i := 1
		for v := range sensor {
			sum += float32(v)
			if i%10 == 0 {
				out <- sum / 10
				sum = 0
			}
			i++
		}
		close(out)
	}()
	return out
}

func CalculateAverage(sum <-chan float32) []float32 {
	res := make([]float32, 0, 6)
	for v := range sum {
		res = append(res, v)
		fmt.Println("Sum is:", v)
	}
	return res
}

func main() {
	sensor := SensorGeneratorChannel()
	sum := AverageSumChannel(sensor)
	CalculateAverage(sum)
}
