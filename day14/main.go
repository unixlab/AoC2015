package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Reindeer struct
type Reindeer struct {
	Name      string
	Speed     int
	RestAfter int
	RestFor   int
	StateFor  int
	Distance  int
	Points    int
	State     bool
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile("([A-Za-z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.")

	var reindeers []Reindeer

	for scanner.Scan() {
		line := scanner.Text()
		regexResult := regex.FindStringSubmatch(line)
		name := regexResult[1]
		speed, _ := strconv.Atoi(regexResult[2])
		restAfter, _ := strconv.Atoi(regexResult[3])
		restFor, _ := strconv.Atoi(regexResult[4])
		reindeer := Reindeer{name, speed, restAfter, restFor, 0, 0, 0, true}
		reindeers = append(reindeers, reindeer)
	}

	for i := 0; i < 2503; i++ {
		for id, reindeer := range reindeers {
			newReindeer := reindeer
			if reindeer.State {
				if reindeer.StateFor == reindeer.RestAfter {
					newReindeer.State = false
					newReindeer.StateFor = 1
				} else {
					newReindeer.StateFor++
					newReindeer.Distance += reindeer.Speed
				}
			} else {
				if reindeer.StateFor == reindeer.RestFor {
					newReindeer.State = true
					newReindeer.StateFor = 1
					newReindeer.Distance += reindeer.Speed
				} else {
					newReindeer.StateFor++
				}
			}
			reindeers[id] = newReindeer
		}
		maxDistance := 0
		for _, reindeer := range reindeers {
			if reindeer.Distance > maxDistance {
				maxDistance = reindeer.Distance
			}
		}
		for id, reindeer := range reindeers {
			if reindeer.Distance == maxDistance {
				newReindeer := reindeer
				newReindeer.Points++
				reindeers[id] = newReindeer
			}
		}
	}

	maxDistance := 0
	for _, reindeer := range reindeers {
		if reindeer.Distance > maxDistance {
			maxDistance = reindeer.Distance
		}
	}
	fmt.Println(maxDistance)

	maxPoints := 0
	for _, reindeer := range reindeers {
		if reindeer.Points > maxPoints {
			maxPoints = reindeer.Points
		}
	}
	fmt.Println(maxPoints)
}
