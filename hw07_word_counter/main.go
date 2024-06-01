package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWords(s string) map[string]int {
	s = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9 ]+`).ReplaceAllString(s, "")
	s = strings.ToLower(s)
	words := strings.Fields(s)
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}
	return wordsMap
}

func main() {
	s := " Hello, World! Hello, Golang  "
	wordsMap := countWords(s)
	for k, v := range wordsMap {
		fmt.Println(k, v)
	}
}
