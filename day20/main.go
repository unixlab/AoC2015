package main

import "fmt"

func main() {
	number := 34000000
	tnumber := number / 10
	housesP1 := make([]int, tnumber)
	housesP2 := make([]int, tnumber)
	for i := 1; i < tnumber; i++ {
		c := 0
		for j := i; j < tnumber; j += i {
			housesP1[j] += i * 10
			c++
			if c < 50 {
				housesP2[j] += i * 11
			}
		}
	}

	for k, v := range housesP1 {
		if v > number {
			fmt.Println(k)
			break
		}
	}

	for k, v := range housesP2 {
		if v > number {
			fmt.Println(k)
			break
		}
	}
}
