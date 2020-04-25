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
	var grid [1000][1000]bool

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		action, xS, yS, xD, yD := parseLine(scanner.Text())
		for x := xS; x <= xD; x++ {
			for y := yS; y <= yD; y++ {
				switch action {
				case 0:
					grid[x][y] = false
				case 1:
					grid[x][y] = true
				case 2:
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	counter := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
