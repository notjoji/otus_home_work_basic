package main

import (
	"fmt"
	"time"
)

func SensorGeneratorChannel(s int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < s; i++ {
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
	outerFor:
		for {
			var sum float32
			for i := 0; i < 10; i++ {
				s, ok := <-sensor
				if !ok {
					break outerFor
				}
				sum += float32(s)
			}
			out <- sum / 10
		}
		close(out)
	}()
	return out
}

func CalculateAverage(sum <-chan float32) []float32 {
	var res []float32
	for {
		s, ok := <-sum
		if !ok {
			break
		}
		res = append(res, s)
		fmt.Println("Sum is:", s)
	}
	return res
}

func main() {
	sensor := SensorGeneratorChannel(60)
	sum := AverageSumChannel(sensor)
	CalculateAverage(sum)
}
