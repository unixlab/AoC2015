package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Cookie struct
type Cookie struct {
	Ingredients []Ingredient
}

func (c Cookie) getScore() int {
	var capacity, durability, flavor, texture int

	for i := 0; i < len(c.Ingredients); i++ {
		capacity += c.Ingredients[i].getCapacity()
		durability += c.Ingredients[i].getDurability()
		flavor += c.Ingredients[i].getFlavor()
		texture += c.Ingredients[i].getTexture()
	}

	if capacity < 0 {
		capacity = 0
	}

	if durability < 0 {
		durability = 0
	}

	if flavor < 0 {
		flavor = 0
	}

	if texture < 0 {
		texture = 0
	}

	return capacity * durability * flavor * texture
}

func (c Cookie) getCalories() int {
	calories := 0
	for i := 0; i < len(c.Ingredients); i++ {
		calories += c.Ingredients[i].getCalories()
	}
	return calories
}

// Ingredient struct
type Ingredient struct {
	Name       string
	Properties Properties
	Amount     int
}

func (i Ingredient) getCapacity() int {
	return i.Properties.Capacity * i.Amount
}

func (i Ingredient) getDurability() int {
	return i.Properties.Durability * i.Amount
}

func (i Ingredient) getFlavor() int {
	return i.Properties.Flavor * i.Amount
}

func (i Ingredient) getTexture() int {
	return i.Properties.Texture * i.Amount
}

func (i Ingredient) getCalories() int {
	return i.Properties.Calories * i.Amount
}

// Properties struct
type Properties struct {
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func main() {
	var ingredients []Ingredient

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	regex := regexp.MustCompile("([A-Za-z]+): capacity ([\\-0-9]+), durability ([\\-0-9]+), flavor ([\\-0-9]+), texture ([\\-0-9]+), calories ([\\-0-9]+)")

	for scanner.Scan() {
		line := scanner.Text()
		regexMatch := regex.FindStringSubmatch(line)

		var ingredient Ingredient
		ingredient.Name = regexMatch[1]
		ingredient.Properties.Capacity, _ = strconv.Atoi(regexMatch[2])
		ingredient.Properties.Durability, _ = strconv.Atoi(regexMatch[3])
		ingredient.Properties.Flavor, _ = strconv.Atoi(regexMatch[4])
		ingredient.Properties.Texture, _ = strconv.Atoi(regexMatch[5])
		ingredient.Properties.Calories, _ = strconv.Atoi(regexMatch[6])

		ingredients = append(ingredients, ingredient)
	}

	maxScore := 0
	maxScoreWith500Calories := 0

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				var cookie Cookie
				l := 100 - i - j - k

				ingredients[0].Amount = i
				ingredients[1].Amount = j
				ingredients[2].Amount = k
				ingredients[3].Amount = l

				cookie.Ingredients = append(cookie.Ingredients, ingredients...)

				score := cookie.getScore()
				calories := cookie.getCalories()

				if score > maxScore {
					maxScore = score
				}

				if calories == 500 && score > maxScoreWith500Calories {
					maxScoreWith500Calories = score
				}
			}
		}
	}

	fmt.Println(maxScore)
	fmt.Println(maxScoreWith500Calories)
}
