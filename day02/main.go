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

func intArrayTwoMin(input []int) (int, int) {
	minKey := 0
	min1 := input[0]
	for k, v := range input {
		if v < min1 {
			minKey = k
			min1 = v
		}
	}

	input = append(input[:minKey], input[minKey+1:]...)

	min2 := input[0]
	for _, v := range input {
		if v < min2 {
			min2 = v
		}
	}

	return min1, min2
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

func calculateRibbon(length int, width int, height int) int {
	sides := make([]int, 3)
	sides[0] = length
	sides[1] = width
	sides[2] = height

	min1, min2 := intArrayTwoMin(sides)
	wrap := min1*2 + min2*2

	bow := length * width * height

	return wrap + bow
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	sumWrappingPaper := 0
	sumRibbon := 0
	for scanner.Scan() {
		l, w, h := extractSizes(scanner.Text())
		sumWrappingPaper += calculateWrappingPaper(l, w, h)
		sumRibbon += calculateRibbon(l, w, h)
	}
	fmt.Println(sumWrappingPaper)
	fmt.Println(sumRibbon)
}
