package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestCalculateAverage(t *testing.T) {
	sensor := CounterSensorGeneratorChannel()
	sum := AverageSumChannel(sensor)
	average := CalculateAverage(sum)
	assert.Equal(t, len(average), 10)
	term := float32(4.5)
	for i, value := range average {
		assert.Equal(t, value, float32(i*10)+term)
	}
}
