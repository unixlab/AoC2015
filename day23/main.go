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

type Registers map[string]int

func (registers Registers) Run(instructions []string) Registers {
	offset := 0
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
	return registers
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	registers := make(Registers, 2)

	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	registers = registers.Run(instructions)

	fmt.Printf("part 1 => %d\n", registers["b"])

	registers["a"] = 1
	registers["b"] = 0

	registers = registers.Run(instructions)

	fmt.Printf("part 2 => %d\n", registers["b"])
}
