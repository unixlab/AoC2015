package main

import (
	"bufio"
	"fmt"
	"os"
)

type House struct {
	X int
	Y int
}

func main() {
	var moves string
	houses := make(map[House]int, 1000)
	x := 0
	y := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves = scanner.Text()
	}

	houses[House{x, y}]++

	for {
		if len(moves) == 0 {
			break
		}

		if moves[0:1] == ">" {
			x++
		} else if moves[0:1] == "<" {
			x--
		} else if moves[0:1] == "^" {
			y++
		} else if moves[0:1] == "v" {
			y--
		}

		houses[House{x, y}]++

		if len(moves) > 0 {
			moves = moves[1:len(moves)]
		} else {
			moves = ""
		}
	}

	fmt.Println(len(houses))
}
