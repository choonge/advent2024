package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	answer1 := 0
	answer2 := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	listLeft := make([]int, 1000)
	listRight := make([]int, 1000)
	mapLeft := make(map[int]int)
	mapRight := make(map[int]int)

	for i := 0; scanner.Scan(); i++ {
		fields := strings.Fields(scanner.Text())
		fieldLeft, _ := strconv.Atoi(fields[0])
		fieldRight, _ := strconv.Atoi(fields[1])

		listLeft[i] = fieldLeft
		listRight[i] = fieldRight

		mapLeft[fieldLeft]++
		mapRight[fieldRight]++
	}

	// Part one
	sort.Ints(listLeft)
	sort.Ints(listRight)

	for k, v := range listLeft {
		distance := v - listRight[k]

		if distance < 0 {
			distance = -distance
		}

		answer1 += distance
	}

	fmt.Printf("Answer part 1: %d\n", answer1)

	// Part two
	for k := range mapLeft {
		answer2 += k * mapRight[k]
	}

	fmt.Printf("Answer part 2: %d\n", answer2)
}
