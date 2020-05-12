package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func combinations(set []int) (subsets [][]int) {
	length := uint(len(set))
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

func intArraySum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}

func main() {
	var numbers []int

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	combinations := combinations(numbers)
	counter := 0
	for _, v := range combinations {
		if intArraySum(v) == 150 {
			counter++
		}
	}
	fmt.Println(counter)
}
