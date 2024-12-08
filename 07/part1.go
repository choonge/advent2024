package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var answer = 0
var operators = []rune{'+', '*'}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		parts := strings.Split(scanner.Text(), ":")
		expected, _ := strconv.Atoi(parts[0])

		numbers := []int{}
		for _, v := range strings.Fields(parts[1]) {
			number, _ := strconv.Atoi(v)
			numbers = append(numbers, number)
		}

		if isValid(numbers, 0, '+', numbers[0], expected, fmt.Sprintf("%d", numbers[0])) || isValid(numbers, 0, '*', numbers[0], expected, fmt.Sprintf("%d", numbers[0])) {
			answer += expected
		}
	}

	// 10741447295096 too high
	fmt.Printf("Answer part 1: %d\n", answer)
}

func isValid(numbers []int, index int, operator rune, result int, expected int, debug string) bool {
	if index == len(numbers)-1 {
		return false
	}

	if operator == '+' {
		result += numbers[index+1]
		debug = fmt.Sprintf("%s+%d", debug, numbers[index+1])
	} else {
		result *= numbers[index+1]
		debug = fmt.Sprintf("%s*%d", debug, numbers[index+1])
	}

	if result > expected {
		return false
	}

	if result == expected {
		fmt.Printf("Debug: %s = %d\n", debug, expected)
		return true
	}

	return isValid(numbers, index+1, '+', result, expected, debug) || isValid(numbers, index+1, '*', result, expected, debug)
}
