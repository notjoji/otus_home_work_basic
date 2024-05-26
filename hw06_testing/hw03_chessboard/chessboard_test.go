package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceAndAtoi(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		error  string
		result int
	}{
		{
			"positive_case",
			"12",
			"",
			12,
		},
		{
			"positive_case_with_sign",
			"-123",
			"",
			-123,
		},
		{
			"not_a_number",
			"asd",
			"strconv.Atoi: parsing \"asd\": invalid syntax",
			0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected, err := ReplaceAndAtoi(testCase.input)
			if err != nil {
				assert.Equal(t, testCase.error, err.Error())
			} else {
				assert.Equal(t, testCase.result, expected)
			}
		})
	}
}

func TestCreateChessboard(t *testing.T) {
	testCases := []struct {
		name   string
		size   int
		result string
	}{
		{
			"size_3",
			3,
			" # \n# #\n # \n",
		},
		{
			"size_5",
			5,
			" # # \n# # #\n # # \n# # #\n # # \n",
		},
		{
			"size_8",
			8,
			" # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected := CreateChessboard(testCase.size)
			assert.Equal(t, testCase.result, expected)
		})
	}
}
