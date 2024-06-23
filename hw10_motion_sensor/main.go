package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func RandomSensorGeneratorChannel(timeout int) chan int {
	c := make(chan int)
	to := time.After(time.Duration(timeout) * time.Second)
	go func() {
		for {
			s, _ := rand.Int(rand.Reader, big.NewInt(100))
			r := int(s.Int64())
			select {
			case <-to:
				close(c)
				return
			case c <- r:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return c
}

func CounterSensorGeneratorChannel() chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
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
	sensor := RandomSensorGeneratorChannel(60)
	sum := AverageSumChannel(sensor)
	CalculateAverage(sum)
}
