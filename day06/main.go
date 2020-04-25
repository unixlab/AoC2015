package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (int, int, int, int, int) {
	var result [5]int

	if strings.HasPrefix(line, "turn on") {
		line = strings.TrimPrefix(line, "turn on ")
		result[0] = 1
	}

	if strings.HasPrefix(line, "turn off") {
		line = strings.TrimPrefix(line, "turn off ")
		result[0] = 0
	}

	if strings.HasPrefix(line, "toggle") {
		line = strings.TrimPrefix(line, "toggle ")
		result[0] = 2
	}

	firstCoordinate := strings.Split(line, " ")[0]
	secondCoordinate := strings.Split(line, " ")[2]

	result[1], _ = strconv.Atoi(strings.Split(firstCoordinate, ",")[0])
	result[2], _ = strconv.Atoi(strings.Split(firstCoordinate, ",")[1])
	result[3], _ = strconv.Atoi(strings.Split(secondCoordinate, ",")[0])
	result[4], _ = strconv.Atoi(strings.Split(secondCoordinate, ",")[1])

	return result[0], result[1], result[2], result[3], result[4]
}

func main() {
	var gridPart1 [1000][1000]bool
	var gridPart2 [1000][1000]int

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		action, xS, yS, xD, yD := parseLine(scanner.Text())
		for x := xS; x <= xD; x++ {
			for y := yS; y <= yD; y++ {
				switch action {
				case 0:
					gridPart1[x][y] = false
					if gridPart2[x][y] > 0 {
						gridPart2[x][y]--
					}
				case 1:
					gridPart1[x][y] = true
					gridPart2[x][y]++
				case 2:
					gridPart1[x][y] = !gridPart1[x][y]
					gridPart2[x][y] += 2
				}
			}
		}
	}

	counterPart1 := 0
	counterPart2 := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if gridPart1[x][y] {
				counterPart1++
			}
			counterPart2 += gridPart2[x][y]
		}
	}

	fmt.Println(counterPart1)
	fmt.Println(counterPart2)
}
