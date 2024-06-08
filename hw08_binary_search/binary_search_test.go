package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{
			"positive case",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			3,
			2,
		},
		{
			"negative case",
			[]int{1, 2, 4, 5, 6, 7, 8, 9, 10},
			3,
			-1,
		},
		{
			"empty array",
			[]int{},
			3,
			-1,
		},
		{
			"one element array",
			[]int{3},
			3,
			0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := BinarySearch(testCase.array, testCase.target)
			assert.Equal(t, testCase.want, got)
		})
	}
}
