package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// change this to run part2
	part2 := true

	wires := make(map[string]uint16, 200)

	regexInput := regexp.MustCompile("^([a-z0-9]+) -> ([a-z]+)$")
	regexAND := regexp.MustCompile("^([a-z0-9]+) AND ([a-z]+) -> ([a-z]+)$")
	regexOR := regexp.MustCompile("^([a-z]+) OR ([a-z]+) -> ([a-z]+)$")
	regexLS := regexp.MustCompile("^([a-z]+) LSHIFT ([0-9]+) -> ([a-z]+)$")
	regexRS := regexp.MustCompile("^([a-z]+) RSHIFT ([0-9]+) -> ([a-z]+)$")
	regexNOT := regexp.MustCompile("^NOT ([a-z]+) -> ([a-z]+)$")

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for len(lines) > 0 {
		line := lines[0]
		lines = lines[1:]
		if regexInput.MatchString(line) {
			tempRegexMatch := regexInput.FindStringSubmatch(line)
			tempInput, err := strconv.Atoi(tempRegexMatch[1])
			var input uint16
			var inMap bool
			if err != nil {
				input, inMap = wires[tempRegexMatch[1]]
			} else {
				input = uint16(tempInput)
				inMap = true
			}
			if !inMap {
				lines = append(lines, line)
				continue
			}
			wire := tempRegexMatch[2]
			if wire == "b" && part2 {
				wires["b"] = 16076
			} else {
				wires[wire] = input
			}
		} else if regexAND.MatchString(line) {
			tempRegexMatch := regexAND.FindStringSubmatch(line)
			wireSrc, err := strconv.Atoi(tempRegexMatch[1])
			var wireSrc1, wireSrc2 uint16
			var inMap1, inMap2 bool
			if err != nil {
				wireSrc1, inMap1 = wires[tempRegexMatch[1]]
			} else {
				wireSrc1 = uint16(wireSrc)
				inMap1 = true
			}
			wireSrc, err = strconv.Atoi(tempRegexMatch[2])
			if err != nil {
				wireSrc2, inMap2 = wires[tempRegexMatch[2]]
			} else {
				wireSrc2 = uint16(wireSrc)
				inMap2 = true
			}
			if !inMap1 || !inMap2 {
				lines = append(lines, line)
				continue
			}
			wireDst := tempRegexMatch[3]
			wires[wireDst] = wireSrc1 & wireSrc2
		} else if regexOR.MatchString(line) {
			tempRegexMatch := regexOR.FindStringSubmatch(line)
			wireSrc1 := tempRegexMatch[1]
			wireSrc2 := tempRegexMatch[2]
			wireDst := tempRegexMatch[3]
			wire1, inMap1 := wires[wireSrc1]
			wire2, inMap2 := wires[wireSrc2]
			if !inMap1 || !inMap2 {
				lines = append(lines, line)
				continue
			}
			wires[wireDst] = wire1 | wire2
		} else if regexLS.MatchString(line) {
			tempRegexMatch := regexLS.FindStringSubmatch(line)
			wireSrc := tempRegexMatch[1]
			shiftBy, _ := strconv.Atoi(tempRegexMatch[2])
			wireDst := tempRegexMatch[3]
			wire, inMap := wires[wireSrc]
			if !inMap {
				lines = append(lines, line)
				continue
			}
			wires[wireDst] = wire << shiftBy
		} else if regexRS.MatchString(line) {
			tempRegexMatch := regexRS.FindStringSubmatch(line)
			wireSrc := tempRegexMatch[1]
			shiftBy, _ := strconv.Atoi(tempRegexMatch[2])
			wireDst := tempRegexMatch[3]
			wire, inMap := wires[wireSrc]
			if !inMap {
				lines = append(lines, line)
				continue
			}
			wires[wireDst] = wire >> shiftBy
		} else if regexNOT.MatchString(line) {
			tempRegexMatch := regexNOT.FindStringSubmatch(line)
			wireSrc := tempRegexMatch[1]
			wireDst := tempRegexMatch[2]
			wire, inMap := wires[wireSrc]
			if !inMap {
				lines = append(lines, line)
				continue
			}
			wires[wireDst] = ^wire
		} else {
			panic(line)
		}
	}
	fmt.Println(wires["a"])
}
