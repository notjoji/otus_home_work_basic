package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	testCases := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			name: "simple positive test",
			text: "  hello   world!   ",
			want: map[string]int{"hello": 1, "world": 1},
		},
		{
			name: "cyrillic test",
			text: "  Привет,  мир!  Привет, язык Go  ",
			want: map[string]int{"привет": 2, "мир": 1, "язык": 1, "go": 1},
		},
		{
			name: "text with numbers test",
			text: "Я напишу 4 теста. Ну какой я молодец!",
			want: map[string]int{"я": 2, "напишу": 1, "4": 1, "теста": 1, "ну": 1, "какой": 1, "молодец": 1},
		},
		{
			name: "empty string test",
			text: "",
			want: map[string]int{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := countWords(testCase.text)
			for k, v := range got {
				assert.Equal(t, v, testCase.want[k], k)
			}
		})
	}
}
