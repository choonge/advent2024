package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer1 := 0

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
		valid := true
		items := strings.Split(line, ",")
	out:
		for j := len(items) - 1; j > 0; j-- {
			for k := 0; k < j; k++ {
				if _, ok := rules[items[j]][items[k]]; ok {
					valid = false
					break out
				}
			}
		}

		if !valid {
			continue
		}

		middle := items[(len(items) / 2)]
		middleInt, _ := strconv.Atoi(middle)
		answer1 += middleInt
	}

	fmt.Printf("Answer part 1: %d\n", answer1)
}
