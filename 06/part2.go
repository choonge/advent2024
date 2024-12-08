package main

import (
	"bufio"
	"fmt"
	"os"
)

var loops = 0

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	grid := make(map[int]map[int]rune)

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

	var restoreY *int
	var restoreX *int

	for y := 0; y <= len(grid); y++ {
		for x := 0; x <= len(grid[0]); x++ {
			if restoreY != nil {
				grid[*restoreY][*restoreX] = '.'
			}

			if grid[y][x] == '.' {
				grid[y][x] = '#'

				restoreY = &y
				restoreX = &x
			}

			visited := make(map[string]int)
			previousVisitedCount := 0

			travel('Y', -1, posY, posX, grid, visited, previousVisitedCount, 0)
		}
	}

	// 1700 = too high
	// 1699 = too high
	fmt.Printf("Answer part 2: %d\n", loops)
}

func travel(axis rune, step int, posY int, posX int, grid map[int]map[int]rune, visited map[string]int, previousVisitedCount int, iteration int) {
	iteration++
	turnedWithoutMoving := false
	previousVisitedCount = len(visited)

	// Travel up
	if axis == 'Y' && step == -1 {
		for i := posY; i >= 0; i-- {
			if grid[i][posX] == '#' {
				if i+1 == posY {
					turnedWithoutMoving = true
				}

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
		for i := posX; i >= 0; i++ {
			if grid[posY][i] == '#' {
				if i-1 == posX {
					turnedWithoutMoving = true
				}

				posX = i - 1
				break
			}

			visited[fmt.Sprintf("%d-%d", posY, i)] = 1

			if i == len(grid[posY])-1 {
				return
			}
		}
	}

	// Travel down
	if axis == 'Y' && step == 1 {
		for i := posY; i >= 0; i++ {
			if grid[i][posX] == '#' {
				if i-1 == posY {
					turnedWithoutMoving = true
				}

				posY = i - 1
				break
			}

			visited[fmt.Sprintf("%d-%d", i, posX)] = 1

			if i == len(grid)+1 {
				return
			}
		}
	}

	// Travel right
	if axis == 'X' && step == -1 {
		for i := posX; i >= 0; i-- {
			if grid[posY][i] == '#' {
				if i+1 == posX {
					turnedWithoutMoving = true
				}

				posX = i + 1
				break
			}

			visited[fmt.Sprintf("%d-%d", posY, i)] = 1

			if i == 0 {
				return
			}
		}
	}

	if previousVisitedCount == len(visited) && !turnedWithoutMoving {
		// We're in a loop
		// loops++
		// return
	}

	if iteration > 1000 {
		loops++
		return
	}

	switch {
	case axis == 'Y' && step == -1:
		// We were going up: go right
		travel('X', 1, posY, posX, grid, visited, previousVisitedCount, iteration)
	case axis == 'X' && step == 1:
		// We were going right: go down
		travel('Y', 1, posY, posX, grid, visited, previousVisitedCount, iteration)
	case axis == 'Y' && step == 1:
		// We were going down: go left
		travel('X', -1, posY, posX, grid, visited, previousVisitedCount, iteration)
	case axis == 'X' && step == -1:
		// We were going left: go up
		travel('Y', -1, posY, posX, grid, visited, previousVisitedCount, iteration)
	}
}
