package main

import "fmt"

func main() {
	number := 34000000
	tnumber := number / 10
	houses := make([]int, tnumber)
	for i := 1; i < tnumber; i++ {
		for j := i; j < tnumber; j += i {
			houses[j] += i * 10
		}
	}

	for k, v := range houses {
		if v > number {
			fmt.Println(k)
			break
		}
	}
}
