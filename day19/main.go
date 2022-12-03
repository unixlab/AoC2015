package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Molecules []string

func (m Molecules) pushUniq(new string) Molecules {
	found := false
	for _, v := range m {
		if v == new {
			found = true
		}
	}
	if !found {
		m = append(m, new)
	}
	return m
}

type Replacement struct {
	From string
	To   string
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	start := lines[len(lines)-1]

	var replacements []Replacement
	for i := 0; i < len(lines)-2; i++ {
		split := strings.Split(lines[i], " => ")
		replacements = append(replacements, Replacement{split[0], split[1]})
	}

	var molecules Molecules
	for _, replacement := range replacements {
		offset := len(replacement.From)
		for i := 0; i <= len(start)-offset; i++ {
			if start[i:i+offset] == replacement.From {
				molecule := fmt.Sprintf("%s%s%s", start[:i], replacement.To, start[i+offset:])
				molecules = molecules.pushUniq(molecule)
				i += offset - 1
			}
		}
	}

	fmt.Println(len(molecules))
}
