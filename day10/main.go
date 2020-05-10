package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1321131112"
	for i := 0; i < 50; i++ {
		var result strings.Builder
		lastCounter := 1
		lastChar := rune(input[0])
		for j := 1; j < len(input); j++ {
			if lastChar == rune(input[j]) {
				lastCounter++
			} else {
				result.WriteString(strconv.Itoa(lastCounter))
				result.WriteRune(lastChar)
				lastCounter = 1
				lastChar = rune(input[j])
			}
		}
		result.WriteString(strconv.Itoa(lastCounter))
		result.WriteRune(lastChar)
		input = result.String()
		if i == 39 {
			fmt.Println(len(input))
		}
	}
	fmt.Println(len(input))
}
