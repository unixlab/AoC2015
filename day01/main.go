package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// we know that we only get one line
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	currentLevel := 0

	for _, moveChar := range strings.Split(line, "") {
		if moveChar == "(" {
			currentLevel++
		} else if moveChar == ")" {
			currentLevel--
		}
	}

	fmt.Println(currentLevel)
}
