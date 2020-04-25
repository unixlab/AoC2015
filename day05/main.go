package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkVowels(input string) bool {
	count := 0
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range vowels {
		count += strings.Count(input, v)
	}
	if count >= 3 {
		return true
	} else {
		return false
	}

}

func checkRow(input string) bool {
	for i := 0; i < len(input); i++ {
		doubleChar := fmt.Sprintf("%s%s", input[i:i+1], input[i:i+1])
		if strings.Index(input, doubleChar) != -1 {
			return true
		}
	}
	return false
}

func checkBlacklist(input string) bool {
	blacklist := []string{"ab", "cd", "pq", "xy"}
	for _, v := range blacklist {
		if strings.Index(input, v) != -1 {
			return false
		}
	}
	return true
}

func main() {
	counter := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		in := scanner.Text()
		if checkBlacklist(in) && checkRow(in) && checkVowels(in) {
			counter++
		}
	}

	fmt.Println(counter)
}
