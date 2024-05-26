package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructComparator(t *testing.T) {
	testCases := []struct {
		name                                      string
		year1, year2                              int32
		size1, size2                              int32
		rate1, rate2                              float32
		fstIsNewer, fstIsBigger, fstHasBetterRate bool
	}{
		{
			"first_is_only_newer",
			1997, 1954,
			694, 1077,
			9.1, 9.3,
			true, false, false,
		},
		{
			"equal_books",
			1997, 1997,
			694, 694,
			9.1, 9.1,
			false, false, false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			book1 := Book{
				1,
				"book1",
				"author1",
				testCase.year1,
				testCase.size1,
				testCase.rate1,
			}
			book2 := Book{
				2,
				"book2",
				"author2",
				testCase.year2,
				testCase.size2,
				testCase.rate2,
			}
			assert.Equal(t, testCase.fstIsNewer, NewBookComparator(Year).Compare(&book1, &book2))
			assert.Equal(t, testCase.fstIsBigger, NewBookComparator(Size).Compare(&book1, &book2))
			assert.Equal(t, testCase.fstHasBetterRate, NewBookComparator(Rate).Compare(&book1, &book2))
		})
	}
}
