package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sue := make(map[string]string, 10)
	sue["children"] = "3"
	sue["cats"] = "7"
	sue["samoyeds"] = "2"
	sue["pomeranians"] = "3"
	sue["akitas"] = "0"
	sue["vizslas"] = "0"
	sue["goldfish"] = "5"
	sue["trees"] = "3"
	sue["cars"] = "2"
	sue["perfumes"] = "1"

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var idPart1, idPart2 string

	for scanner.Scan() {
		line := strings.TrimPrefix(scanner.Text(), "Sue ")
		line = strings.ReplaceAll(line, " ", "")
		firstColon := strings.Index(line, ":")
		sueID := line[0:firstColon]
		line = line[firstColon+1:]
		compounds := strings.Split(line, ",")

		match := true
		for _, compound := range compounds {
			c := strings.Split(compound, ":")
			if sue[c[0]] != c[1] {
				match = false
			}
		}

		if match {
			idPart1 = sueID
		}

		match = true

		for _, compound := range compounds {
			c := strings.Split(compound, ":")
			if c[0] == "cats" || c[0] == "trees" {
				c0, _ := strconv.Atoi(sue[c[0]])
				c1, _ := strconv.Atoi(c[1])
				if c0 >= c1 {
					match = false
				}
			} else if c[0] == "pomeranians" || c[0] == "goldfish" {
				c0, _ := strconv.Atoi(sue[c[0]])
				c1, _ := strconv.Atoi(c[1])
				if c0 <= c1 {
					match = false
				}
			} else {
				if sue[c[0]] != c[1] {
					match = false
				}
			}

		}

		if match {
			idPart2 = sueID
		}

	}
	fmt.Println(idPart1)
	fmt.Println(idPart2)
}
