package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func getIntSum(value interface{}) int {
	sum := 0
	switch value.(type) {
	case float64:
		sum += int(value.(float64))
	case []interface{}:
		for _, v := range value.([]interface{}) {
			sum += getIntSum(v)
		}
	case map[string]interface{}:
		for _, v := range value.(map[string]interface{}) {
			sum += getIntSum(v)
		}
	}
	return sum
}

func getIntSumWithoutRed(value interface{}) (int, bool) {
	sum := 0
	switch value.(type) {
	case string:
		if value.(string) == "red" {
			return 0, false
		}
	case float64:
		sum += int(value.(float64))
	case []interface{}:
		for _, v := range value.([]interface{}) {
			if childSum, ok := getIntSumWithoutRed(v); ok {
				sum += childSum
			}
		}
	case map[string]interface{}:
		for _, v := range value.(map[string]interface{}) {
			if childSum, ok := getIntSumWithoutRed(v); ok {
				sum += childSum
			} else {
				return 0, true
			}
		}
	}
	return sum, true
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var input string
	for scanner.Scan() {
		input = scanner.Text()

		var jsonInterface interface{}
		json.Unmarshal([]byte(input), &jsonInterface)

		fmt.Println(getIntSum(jsonInterface))
		fmt.Println(getIntSumWithoutRed(jsonInterface))
	}
}
