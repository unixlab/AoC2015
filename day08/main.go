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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineUnquoted, _ := strconv.Unquote(line)
		sum += len(line) - len(lineUnquoted)
	}
	fmt.Println(sum)
}
