package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	text, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	text = strings.ReplaceAll(text, "\n", "")
	size, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
		return
	}
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
	fmt.Println(result)
}
