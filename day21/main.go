package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type item struct {
	Cost   int
	Damage int
	Armor  int
}

type fighter struct {
	HP     int
	Damage int
	Armor  int
}

func fight(p1 fighter, p2 fighter) bool {
	for p1.HP > 0 && p2.HP > 0 {
		p2.HP -= p1.Damage - p2.Armor
		if p2.HP <= 0 {
			return true
		}
		p1.HP -= p2.Damage - p1.Armor
		if p1.HP <= 0 {
			return false
		}
	}
	panic("error in fight")
}

func main() {
	var weapons [5]item
	var armors [6]item
	var rings [8]item

	weapons[0] = item{8, 4, 0}
	weapons[1] = item{10, 5, 0}
	weapons[2] = item{25, 6, 0}
	weapons[3] = item{40, 7, 0}
	weapons[4] = item{74, 8, 0}
	armors[0] = item{13, 0, 1}
	armors[1] = item{31, 0, 2}
	armors[2] = item{53, 0, 3}
	armors[3] = item{75, 0, 4}
	armors[4] = item{102, 0, 5}
	rings[0] = item{25, 1, 0}
	rings[1] = item{50, 2, 0}
	rings[2] = item{100, 3, 0}
	rings[3] = item{20, 0, 1}
	rings[4] = item{40, 0, 2}
	rings[5] = item{80, 0, 3}

	var boss fighter

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Index(line, ":")
		number, _ := strconv.Atoi(line[pos+2:])
		if strings.HasPrefix(line, "Hit Points") {
			boss.HP = number
		} else if strings.HasPrefix(line, "Damage") {
			boss.Damage = number
		} else if strings.HasPrefix(line, "Armor") {
			boss.Armor = number
		} else {
			fmt.Printf("error in input file %s\n", line)
		}
	}

	var you fighter
	minCosts := 10000

	for _, weapon := range weapons {
		for _, armor := range armors {
			ringList := combin.Combinations(8, 2)
			for _, ring := range ringList {
				costs := weapon.Cost + armor.Cost + rings[ring[0]].Cost + rings[ring[1]].Cost
				you.HP = 100
				you.Armor = armor.Armor + rings[ring[0]].Armor + rings[ring[1]].Armor
				you.Damage = weapon.Damage + rings[ring[0]].Damage + rings[ring[1]].Damage
				if fight(you, boss) && costs < minCosts {
					minCosts = costs
				}
			}
		}
	}

	fmt.Println(minCosts)
}
