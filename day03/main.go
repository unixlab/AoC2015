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
	// part1
	var moves string
	houses := make(map[House]int, 1000)
	x := 0
	y := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves = scanner.Text()
	}
	file.Close()

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

	// part2
	houses = make(map[House]int, 1000)
	xS := 0
	yS := 0
	xR := 0
	yR := 0

	file, _ = os.Open("input.txt")
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		moves = scanner.Text()
	}
	file.Close()

	houses[House{xS, yS}]++
	houses[House{xR, yR}]++

	for {
		if len(moves) == 0 {
			break
		}

		if moves[0:1] == ">" {
			xS++
		} else if moves[0:1] == "<" {
			xS--
		} else if moves[0:1] == "^" {
			yS++
		} else if moves[0:1] == "v" {
			yS--
		}

		if moves[1:2] == ">" {
			xR++
		} else if moves[1:2] == "<" {
			xR--
		} else if moves[1:2] == "^" {
			yR++
		} else if moves[1:2] == "v" {
			yR--
		}

		houses[House{xS, yS}]++
		houses[House{xR, yR}]++

		if len(moves) > 0 {
			moves = moves[2:len(moves)]
		} else {
			moves = ""
		}
	}

	fmt.Println(len(houses))
}
