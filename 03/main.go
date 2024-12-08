package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	answer1 := 0
	answer2 := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	enabled := true

	for i := 0; scanner.Scan(); i++ {
		re := regexp.MustCompile(`(?:don\'t\(\))|(?:do\(\))|(?:mul\((?P<first>\d+),(?P<second>\d+)\))`)
		matches := re.FindAllStringSubmatch(scanner.Text(), 99999)

		for _, v := range matches {
			fmt.Printf("%#v\n", v)

			if v[0] == "do()" {
				enabled = true
			} else if (v[0]) == "don't()" {
				enabled = false
			}

			if v[1] != "" {
				first, _ := strconv.Atoi(v[1])
				second, _ := strconv.Atoi(v[2])

				// Part one
				answer1 += first * second

				// Part two
				// 91634027 too high
				// 89349241
				if enabled == true {
					answer2 += first * second
					fmt.Printf("Enabled - %d * %d (%d)\n\n", first, second, answer2)
				}
			}
		}
	}

	fmt.Printf("Answer part 1: %d\n", answer1)
	fmt.Printf("Answer part 2: %d\n", answer2)
}
