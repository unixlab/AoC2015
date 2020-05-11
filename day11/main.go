package main

import (
	"fmt"
	"strings"
)

func checkPassword(password string) bool {
	// rule 2
	if strings.Index(password, "i") >= 0 {
		return false
	}
	if strings.Index(password, "o") >= 0 {
		return false
	}
	if strings.Index(password, "l") >= 0 {
		return false
	}
	// rule 3
	overlap := 0
	for i := 97; i <= 122; i++ {
		if strings.Index(password, fmt.Sprintf("%c%c", i, i)) >= 0 {
			overlap++
		}
	}
	if overlap < 2 {
		return false
	}
	// rule 1
	found := false
	for i := 97; i <= 120; i++ {
		if strings.Index(password, fmt.Sprintf("%c%c%c", i, i+1, i+2)) >= 0 {
			found = true
		}
	}
	if !found {
		return false
	}
	// all good
	return true
}

func incrementPassword(input string) string {
	var passwordAsInt [8]int
	for i := 0; i < 8; i++ {
		passwordAsInt[i] = int(input[i])
	}
	flipped := false
	for i := 7; i >= 0; i-- {
		passwordAsInt[i]++
		if passwordAsInt[i] > 122 {
			passwordAsInt[i] = 97
			flipped = true
		} else {
			flipped = false
		}
		if !flipped {
			break
		}
	}
	var newInput strings.Builder
	for i := 0; i < 8; i++ {
		newInput.WriteRune(rune(passwordAsInt[i]))
	}
	return newInput.String()
}

func main() {
	input := "vzbxkghb"
	for !checkPassword(input) {
		input = incrementPassword(input)
	}
	fmt.Println(input)
}
