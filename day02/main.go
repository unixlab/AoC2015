package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intArrayMin(input []int) int {
	min := input[0]
	for _, v := range input {
		if v < min {
			min = v
		}
	}
	return min
}

func calculateWrappingPaper(length int, width int, height int) int {
	sides := make([]int, 3)
	sides[0] = length * width
	sides[1] = width * height
	sides[2] = height * length

	extra := intArrayMin(sides)

	return 2*sides[0] + 2*sides[1] + 2*sides[2] + extra
}

func extractSizes(input string) (int, int, int) {
	numbersAsString := strings.Split(input, "x")

	length, _ := strconv.Atoi(numbersAsString[0])
	width, _ := strconv.Atoi(numbersAsString[1])
	height, _ := strconv.Atoi(numbersAsString[2])

	return length, width, height
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += calculateWrappingPaper(extractSizes(scanner.Text()))
	}
	fmt.Println(sum)
}
