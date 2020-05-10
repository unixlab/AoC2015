package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	sumPart1 := 0
	sumPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineUnquoted, _ := strconv.Unquote(line)
		lineQuoted := strconv.Quote(line)
		sumPart1 += len(line) - len(lineUnquoted)
		sumPart2 += len(lineQuoted) - len(line)
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}
