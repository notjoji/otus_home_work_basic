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
	assert.Equal(t, average[0], float32(4.5))
	assert.Equal(t, average[1], float32(14.5))
}
