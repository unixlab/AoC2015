package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	prmt "github.com/gitchander/permutation"
)

// Happiness our struct to save happiness
type Happiness struct {
	Person    string
	Happiness int
}

func getHappiness(sittingData map[string][]Happiness, person1 string, person2 string) int {
	for _, v := range sittingData[person1] {
		if v.Person == person2 {
			return v.Happiness
		}
	}
	return 0
}

func main() {
	sittingData := make(map[string][]Happiness, 10)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.ReplaceAll(line, "would ", "")
		line = strings.ReplaceAll(line, "happiness units by sitting next to ", "")
		line = line[:len(line)-1]

		splitted := strings.Split(line, " ")

		gain, _ := strconv.Atoi(splitted[2])

		if splitted[1] == "lose" {
			gain = gain * -1
		}

		sittingData[splitted[0]] = append(sittingData[splitted[0]], Happiness{splitted[3], gain})
	}

	var names []string
	for name := range sittingData {
		names = append(names, name)
	}

	numberOfPersons := len(names)
	bestArrangement := -1

	p := prmt.New(prmt.StringSlice(names))
	for p.Next() {
		currentChange := 0
		for i := 0; i < numberOfPersons; i++ {
			var person1, person2 string
			person1 = names[i]
			if i+1 == numberOfPersons {
				person2 = names[0]
			} else {
				person2 = names[i+1]
			}
			currentChange += getHappiness(sittingData, person1, person2)
			currentChange += getHappiness(sittingData, person2, person1)
		}
		if currentChange > bestArrangement || bestArrangement == -1 {
			bestArrangement = currentChange
		}
	}
	fmt.Println(bestArrangement)

	for name := range sittingData {
		sittingData[name] = append(sittingData[name], Happiness{"me", 0})
		sittingData["me"] = append(sittingData["me"], Happiness{name, 0})
	}

	names = append(names, "me")

	numberOfPersons = len(names)
	bestArrangement = -1

	p = prmt.New(prmt.StringSlice(names))
	for p.Next() {
		currentChange := 0
		for i := 0; i < numberOfPersons; i++ {
			var person1, person2 string
			person1 = names[i]
			if i+1 == numberOfPersons {
				person2 = names[0]
			} else {
				person2 = names[i+1]
			}
			currentChange += getHappiness(sittingData, person1, person2)
			currentChange += getHappiness(sittingData, person2, person1)
		}
		if currentChange > bestArrangement || bestArrangement == -1 {
			bestArrangement = currentChange
		}
	}
	fmt.Println(bestArrangement)
}
