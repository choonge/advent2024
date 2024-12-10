package main

import (
	"bufio"
	"fmt"
	"os"
)

var island = []string{}
var islandHeight int
var islandWidth int

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		island = append(island, scanner.Text())
	}

	islandHeight = len(island)
	islandWidth = len(island[0])

	answer1 := 0
	answer2 := 0

	for y, line := range island {
		for x, char := range line {
			// Find all zeros
			if char != 48 {
				continue
			}

			// Part one
			trailhead := fmt.Sprintf("%d-%d", y, x)
			peaks := make(map[string]bool)
			searchPart1(y, x, trailhead, peaks)
			answer1 += len(peaks)

			// Part two
			var trails *int = new(int)
			*trails = 0
			searchPart2(y, x, trails)
			answer2 += *trails
		}
	}

	fmt.Printf("Answer part 1: %d\n", answer1)
	fmt.Printf("Answer part 2: %d\n", answer2)
}

func searchPart2(y int, x int, trails *int) {
	currentValue := island[y][x]
	searchValue := currentValue + 1

	if currentValue == 57 {
		*trails++
		return
	}

	// Travel above
	if y-1 >= 0 && island[y-1][x] == searchValue {
		searchPart2(y-1, x, trails)
	}

	// Travel right
	if x+1 < islandWidth && island[y][x+1] == searchValue {
		searchPart2(y, x+1, trails)
	}

	// Travel below
	if y+1 < islandHeight && island[y+1][x] == searchValue {
		searchPart2(y+1, x, trails)
	}

	// Travel left
	if x-1 >= 0 && island[y][x-1] == searchValue {
		searchPart2(y, x-1, trails)
	}
}

func searchPart1(y int, x int, trailhead string, peaks map[string]bool) {
	currentValue := island[y][x]
	searchValue := currentValue + 1

	if currentValue == 57 {
		peaks[fmt.Sprintf("%s-%d-%d", trailhead, y, x)] = true
		return
	}

	// Travel above
	if y-1 >= 0 && island[y-1][x] == searchValue {
		searchPart1(y-1, x, trailhead, peaks)
	}

	// Travel right
	if x+1 < islandWidth && island[y][x+1] == searchValue {
		searchPart1(y, x+1, trailhead, peaks)
	}

	// Travel below
	if y+1 < islandHeight && island[y+1][x] == searchValue {
		searchPart1(y+1, x, trailhead, peaks)
	}

	// Travel left
	if x-1 >= 0 && island[y][x-1] == searchValue {
		searchPart1(y, x-1, trailhead, peaks)
	}
}
