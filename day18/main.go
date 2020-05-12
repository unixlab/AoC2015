package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	var grid [102][102]bool

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	x := 1
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i:i+1] == "#" {
				grid[x][i+1] = true
			} else {
				grid[x][i+1] = false
			}
		}
		x++
	}

	for i := 0; i < 100; i++ {
		newGrid := grid

		for x, row := range grid {
			if x == 0 || x == len(grid)-1 {
				continue
			}
			for y, light := range row {
				if y == 0 || y == len(grid)-1 {
					continue
				}
				sumLights := 0
				for a := x - 1; a < x+2; a++ {
					for b := y - 1; b < y+2; b++ {
						if a == x && b == y {
							continue
						}
						if grid[a][b] {
							sumLights++
						}
					}
				}
				if light {
					if sumLights == 2 || sumLights == 3 {
						newGrid[x][y] = true
					} else {
						newGrid[x][y] = false
					}
				} else {
					if sumLights == 3 {
						newGrid[x][y] = true
					} else {
						newGrid[x][y] = false
					}
				}
			}
		}

		grid = newGrid
	}

	counter := 0
	for x, v := range grid {
		if x == 0 || x == len(grid)-1 {
			continue
		}
		for y, light := range v {
			if y == 0 || y == len(grid)-1 {
				continue
			}
			if light {
				counter++
			}
		}
	}
	fmt.Println(counter)
}

func part2() {
	var grid [102][102]bool

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	x := 1
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i:i+1] == "#" {
				grid[x][i+1] = true
			} else {
				grid[x][i+1] = false
			}
		}
		x++
	}

	grid[1][1] = true
	grid[1][100] = true
	grid[100][1] = true
	grid[100][100] = true

	for i := 0; i < 100; i++ {
		newGrid := grid

		for x, row := range grid {
			if x == 0 || x == len(grid)-1 {
				continue
			}
			for y, light := range row {
				if y == 0 || y == len(grid)-1 {
					continue
				}
				if x == 1 && y == 1 || x == 1 && y == 100 || x == 100 && y == 1 || x == 100 && y == 100 {
					continue
				}
				sumLights := 0
				for a := x - 1; a < x+2; a++ {
					for b := y - 1; b < y+2; b++ {
						if a == x && b == y {
							continue
						}
						if grid[a][b] {
							sumLights++
						}
					}
				}
				if light {
					if sumLights == 2 || sumLights == 3 {
						newGrid[x][y] = true
					} else {
						newGrid[x][y] = false
					}
				} else {
					if sumLights == 3 {
						newGrid[x][y] = true
					} else {
						newGrid[x][y] = false
					}
				}
			}
		}

		grid = newGrid
	}

	counter := 0
	for x, v := range grid {
		if x == 0 || x == len(grid)-1 {
			continue
		}
		for y, light := range v {
			if y == 0 || y == len(grid)-1 {
				continue
			}
			if light {
				counter++
			}
		}
	}
	fmt.Println(counter)
}

func main() {
	part1()
	part2()
}
