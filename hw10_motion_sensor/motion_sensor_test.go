package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
	sensor := SensorGeneratorChannel(20)
	sum := AverageSumChannel(sensor)
	average := CalculateAverage(sum)
	assert.Equal(t, len(average), 2)
	term := float32(4.5)
	for i, value := range average {
		assert.Equal(t, value, float32(i*10)+term)
	}
}
