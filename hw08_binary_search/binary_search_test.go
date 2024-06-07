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
		want   bool
	}{
		{
			"positive case",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			3,
			true,
		},
		{
			"negative case",
			[]int{1, 2, 4, 5, 6, 7, 8, 9, 10},
			3,
			false,
		},
		{
			"empty array",
			[]int{},
			3,
			false,
		},
		{
			"one element array",
			[]int{3},
			3,
			true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := BinarySearch(testCase.array, testCase.target)
			assert.Equal(t, testCase.want, got)
		})
	}
}
