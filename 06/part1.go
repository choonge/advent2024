package main

import (
	"bufio"
	"fmt"
	"os"
)

var grid = make(map[int]map[int]rune)
var visited = make(map[string]int)

func main() {
	file, _ := os.Open("input_test.txt")
	scanner := bufio.NewScanner(file)

	posY := 0
	posX := 0

	for y := 0; scanner.Scan(); y++ {
		for x, char := range scanner.Text() {
			if char == '^' {
				posY = y
				posX = x

				continue
			}

			if _, ok := grid[y]; !ok {
				grid[y] = make(map[int]rune)
			}

			grid[y][x] = char
		}
	}

	fmt.Printf("Starting at y %d, x %d, traveling upwards\n", posY, posX)

	travel('Y', -1, posY, posX)

	// 4188 too low
	fmt.Printf("Answer part 1: %d\n", len(visited))
}

func travel(axis rune, step int, posY int, posX int) {
	// Travel up
	if axis == 'Y' && step == -1 {
		for i := posY; i >= 0; i-- {
			fmt.Printf("- At y %d, x %d...\n", i, posX)

			if grid[i][posX] == '#' {
				fmt.Printf("Found # at y %d, x %d\n", i, posX)
				posY = i + 1
				break
			}

			visited[fmt.Sprintf("%d-%d", i, posX)] = 1

			if i == 0 {
				return
			}
		}
	}

	// Travel right
	if axis == 'X' && step == 1 {
		for posX = posX; posX >= 0; posX++ {
			fmt.Printf("- At y %d, x %d...\n", posY, posX)

			if grid[posY][posX] == '#' {
				fmt.Printf("Found # at y %d, x %d\n", posY, posX)
				posX--
				break
			}

			visited[fmt.Sprintf("%d-%d", posY, posX)] = 1

			if posX == len(grid[posY])-1 {
				return
			}
		}
	}

	// Travel down
	if axis == 'Y' && step == 1 {
		for posY = posY; posY >= 0; posY++ {
			fmt.Printf("- At y %d, x %d...\n", posY, posX)

			if grid[posY][posX] == '#' {
				fmt.Printf("Found # at y %d, x %d\n", posY, posX)
				posY--
				break
			}

			visited[fmt.Sprintf("%d-%d", posY, posX)] = 1

			if posY == len(grid)+1 {
				return
			}
		}
	}

	// Travel right
	if axis == 'X' && step == -1 {
		for posX = posX; posX >= 0; posX-- {
			fmt.Printf("- At y %d, x %d...\n", posY, posX)

			if grid[posY][posX] == '#' {
				fmt.Printf("Found # at y %d, x %d\n", posY, posX)
				posX++
				break
			}

			visited[fmt.Sprintf("%d-%d", posY, posX)] = 1

			if posX == 0 {
				return
			}
		}
	}

	switch {
	case axis == 'Y' && step == -1:
		// We were going up: go right
		fmt.Printf("Turning right at y %d, x %d\n", posY, posX)
		travel('X', 1, posY, posX)
	case axis == 'X' && step == 1:
		// We were going right: go down
		fmt.Printf("Turning down at y %d, x %d\n", posY, posX)
		travel('Y', 1, posY, posX)
	case axis == 'Y' && step == 1:
		// We were going down: go left
		fmt.Printf("Turning left at y %d, x %d\n", posY, posX)
		travel('X', -1, posY, posX)
	case axis == 'X' && step == -1:
		// We were going left: go up
		fmt.Printf("Turning up at y %d, x %d\n", posY, posX)
		travel('Y', -1, posY, posX)
	}
}
