package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Replacement struct
type Replacement struct {
	Find    string
	Replace string
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var molecule string
	var replacements []Replacement

	regexReplacements := regexp.MustCompile("^([A-Za-z]+) => ([A-Za-z]+)$")
	regexMolecule := regexp.MustCompile("^[A-Za-z]+$")

	for scanner.Scan() {
		line := scanner.Text()
		if regexReplacements.MatchString(line) {
			tempRegexRes := regexReplacements.FindStringSubmatch(line)
			replacements = append(replacements, Replacement{tempRegexRes[1], tempRegexRes[2]})
		}
		if regexMolecule.MatchString(line) {
			molecule = line
		}
	}

	var replacedString strings.Builder
	molecules := make(map[string]int)

	for _, replacement := range replacements {
		pos := 0
		for strings.Index(molecule[pos:], replacement.Find) > -1 {
			replacedString.Reset()
			pos += strings.Index(molecule[pos:], replacement.Find)
			replacedString.WriteString(molecule[:pos])
			replacedString.WriteString(replacement.Replace)
			pos += len(replacement.Find)
			replacedString.WriteString(molecule[pos:])
			molecules[replacedString.String()]++
		}
	}

	fmt.Println(len(molecules))
}
