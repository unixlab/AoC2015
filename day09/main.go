package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	prmt "github.com/gitchander/permutation"
)

// Distance is our own struct to save distances
type Distance struct {
	City     string
	Distance int
}

func getDistance(citys map[string][]Distance, cityA string, cityB string) int {
	for _, v := range citys[cityA] {
		if v.City == cityB {
			return v.Distance
		}
	}
	return 0
}

func main() {
	citys := make(map[string][]Distance, 10)

	regex := regexp.MustCompile("^([A-Za-z]+) to ([A-Za-z]+) = ([0-9]+)$")

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tempRegexMatch := regex.FindStringSubmatch(scanner.Text())
		cityA := tempRegexMatch[1]
		cityB := tempRegexMatch[2]
		distance, _ := strconv.Atoi(tempRegexMatch[3])

		citys[cityA] = append(citys[cityA], Distance{cityB, distance})
		citys[cityB] = append(citys[cityB], Distance{cityA, distance})
	}

	var cityNames []string
	for name := range citys {
		cityNames = append(cityNames, name)
	}

	minDistance := -1
	maxDistance := -1

	p := prmt.New(prmt.StringSlice(cityNames))
	for p.Next() {
		distance := 0
		cityBefore := ""
		for _, v := range cityNames {
			if cityBefore == "" {
				cityBefore = v
			} else {
				distance += getDistance(citys, cityBefore, v)
				cityBefore = v
			}
		}
		if distance < minDistance || minDistance == -1 {
			minDistance = distance
		}
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	fmt.Println(minDistance)
	fmt.Println(maxDistance)
}
