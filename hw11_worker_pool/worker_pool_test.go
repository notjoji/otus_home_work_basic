package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitWorkerPool(t *testing.T) {
	testCases := []struct {
		name string
		size int
		want int
	}{
		{
			"simple_test",
			1000,
			1000,
		},
		{
			"lots_of_workers_test",
			1000000,
			1000000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, InitWorkerPool(tc.size, false))
		})
	}
}
