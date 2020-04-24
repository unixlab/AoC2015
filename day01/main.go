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
	firstBasementHit := 0

	for move, moveChar := range strings.Split(line, "") {
		if moveChar == "(" {
			currentLevel++
		} else if moveChar == ")" {
			currentLevel--
		}
		if currentLevel == -1 && firstBasementHit == 0 {
			firstBasementHit = move + 1
		}
	}

	fmt.Println(currentLevel)
	fmt.Println(firstBasementHit)
}
