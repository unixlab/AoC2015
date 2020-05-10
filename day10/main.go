package main

import "fmt"

func getNewInput(input string) string {
	return fmt.Sprintf("%d%s", len(input), input[0:1])
}

func main() {
	input := "1321131112"
	for i := 0; i < 40; i++ {
		j := 0
		newInput := ""
		for j < len(input)-1 {
			if input[j] != input[j+1] {
				newInput += getNewInput(input[:j+1])
				input = input[j+1:]
				j = 0
				continue
			}
			j++
		}
		newInput += getNewInput(input)
		input = newInput
	}
	fmt.Println(len(input))
}
