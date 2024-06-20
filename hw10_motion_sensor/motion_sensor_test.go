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
	for _, v := range average {
		assert.GreaterOrEqual(t, v, float32(0))
		assert.LessOrEqual(t, v, float32(100))
	}
}
