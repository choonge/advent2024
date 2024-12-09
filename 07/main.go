package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var answer1 = 0
var answer2 = 0

func main() {
	file, _ := os.Open("input_test.txt")
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		parts := strings.Split(scanner.Text(), ":")
		expected, _ := strconv.Atoi(parts[0])

		numbers := []int{}
		for _, v := range strings.Fields(parts[1]) {
			number, _ := strconv.Atoi(v)
			numbers = append(numbers, number)
		}

		for _, attempt := range []rune{'+', '*'} {
			if isValidPart1(numbers, 0, attempt, numbers[0], expected) {
				answer1 += expected
				break
			}
		}

		for _, attempt := range []string{"+", "*", "||"} {
			if isValidPart2(numbers, 0, attempt, numbers[0], expected) {
				answer2 += expected
				break
			}
		}
	}

	// 10741447295096 too high
	// 10741443549536 correct
	fmt.Printf("Answer part 1: %d\n", answer1)

	// 15916264547402 too low
	fmt.Printf("Answer part 2: %d\n", answer2)
}

func isValidPart2(numbers []int, index int, operator string, result int, expected int) bool {
	if index == len(numbers)-1 {
		return result == expected
	}

	switch operator {
	case "+":
		result += numbers[index+1]
	case "*":
		result *= numbers[index+1]
	case "||":
		x, _ := strconv.Atoi(fmt.Sprintf("%d%d", numbers[index], numbers[index+1]))
		fmt.Printf("%d: %d||%d\n", index, numbers[index], numbers[index+1])
		result = x
	}

	for _, attempt := range []string{"+", "*", "||"} {
		if isValidPart2(numbers, index+1, attempt, result, expected) {
			return true
		}
	}

	return false
}

func isValidPart1(numbers []int, index int, operator rune, result int, expected int) bool {
	if index == len(numbers)-1 {
		return result == expected
	}

	switch operator {
	case '+':
		result += numbers[index+1]
	case '*':
		result *= numbers[index+1]
	}

	if result > expected {
		return false
	}

	for _, attempt := range []rune{'+', '*'} {
		if isValidPart1(numbers, index+1, attempt, result, expected) {
			return true
		}
	}

	return false
}
