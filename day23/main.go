package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToIntWithAlgebraicSign(input string) int {
	var result int
	input = strings.TrimSpace(input)
	if strings.HasPrefix(input, "+") {
		input = strings.Replace(input, "+", "", -1)
		result, _ = strconv.Atoi(input)
	} else {
		input = strings.Replace(input, "+", "", -1)
		result, _ = strconv.Atoi(input)
		result *= -1
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	registers := make(map[string]int, 2)

	offset := 0
	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	for offset < len(instructions) {
		instruction := instructions[offset]
		switch {
		case strings.HasPrefix(instruction, "hlf"):
			reg := strings.Split(instruction, " ")[1]
			registers[reg] = registers[reg] / 2
			offset++
		case strings.HasPrefix(instruction, "tpl"):
			reg := strings.Split(instruction, " ")[1]
			registers[reg] = registers[reg] * 3
			offset++
		case strings.HasPrefix(instruction, "inc"):
			reg := strings.Split(instruction, " ")[1]
			registers[reg]++
			offset++
		case strings.HasPrefix(instruction, "jmp"):
			jumpString := strings.Split(instruction, " ")[1]
			jumpInt, _ := strconv.Atoi(jumpString)
			offset += jumpInt
		case strings.HasPrefix(instruction, "jie"):
			reg := strings.Split(instruction, " ")[1][0:1]
			if registers[reg]%2 == 0 {
				jumpString := strings.Split(instruction, ",")[1]
				jumpInt := stringToIntWithAlgebraicSign(jumpString)
				offset += jumpInt
			} else {
				offset++
			}
		case strings.HasPrefix(instruction, "jio"):
			reg := strings.Split(instruction, " ")[1][0:1]
			if registers[reg] == 1 {
				jumpString := strings.Split(instruction, ",")[1]
				jumpInt := stringToIntWithAlgebraicSign(jumpString)
				offset += jumpInt
			} else {
				offset++
			}
		default:
			panic("unknown instruction in input")
		}
	}

	fmt.Printf("part 1 => %d\n", registers["b"])
}
