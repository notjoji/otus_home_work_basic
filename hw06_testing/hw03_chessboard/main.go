package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadSize() (string, error) {
	scanner := bufio.NewReader(os.Stdin)
	text, err := scanner.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("ошибка при вводе: %s", err)
	}
	return text, nil
}

func ReplaceAndAtoi(s string) (int, error) {
	s = strings.ReplaceAll(s, "\n", "")
	return strconv.Atoi(s)
}

func CreateChessboard(size int) string {
	var result string
	for i := 0; i < size; i++ {
		var line string
		for j := 0; j < size; j++ {
			s := " "
			if (i+j)%2 == 1 {
				s = "#"
			}
			line += s
		}
		result += line + "\n"
	}
	return result
}

func main() {
	text, err := ReadSize()
	if err != nil {
		panic(err)
	}
	size, err := ReplaceAndAtoi(text)
	if err != nil {
		panic(err)
	}
	fmt.Println(CreateChessboard(size))
}
