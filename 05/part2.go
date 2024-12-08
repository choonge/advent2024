package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer2 := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	rules := make(map[string]map[string]int)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// Add rule
		if len(line) == 5 {
			parts := strings.Split(line, "|")

			if rules[parts[0]] == nil {
				rules[parts[0]] = make(map[string]int)
			}

			rules[parts[0]][parts[1]] = 1

			continue
		}

		// Check the line
		items := strings.Split(line, ",")
		valid, j, k := isValid(items, rules)

		if valid {
			continue
		}

		// Try to correctly order the line
		for valid == false {
			fmt.Printf("Line incorrect: %#v\n", line)
			fmt.Printf("Swapping offending elements: %#v, %#v\n", j, k)

			first := items[j]
			second := items[k]

			items[j] = second
			items[k] = first

			valid, j, k = isValid(items, rules)

			fmt.Println("")
		}

		middle := items[(len(items) / 2)]
		middleInt, _ := strconv.Atoi(middle)
		answer2 += middleInt
	}

	fmt.Printf("Answer part 2: %d\n", answer2)
}

func isValid(items []string, rules map[string]map[string]int) (bool, int, int) {
	for j := len(items) - 1; j > 0; j-- {
		for k := 0; k < j; k++ {
			if _, ok := rules[items[j]][items[k]]; ok {
				return false, j, k
			}
		}
	}

	return true, 0, 0
}
