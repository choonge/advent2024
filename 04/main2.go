package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	answer1 := 0
	answer2 := 0

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lines := make([][]rune, 140)

	// Read file into memory
	for i := 0; scanner.Scan(); i++ {
		lines[i] = []rune(scanner.Text())
	}

	for row, line := range lines {
		for col, char := range line {
			if char != 'A' {
				continue
			}

			fmt.Printf("Row %d | Col %d | %c\n", row, col, char)

			spaceRight := col+1 < len(line)
			spaceUnder := row+1 < len(lines)
			spaceLeft := col-1 >= 0
			spaceUpper := row-1 >= 0

			topToBottom := ""
			bottomToTop := ""

			// Top to bottom
			if spaceRight && spaceUnder && spaceLeft && spaceUpper {
				topToBottom = fmt.Sprintf("%s%s%s", string(lines[row-1][col-1:col]), string(char), string(lines[row+1][col+1:col+2]))
				fmt.Printf("Top to bottom: %s\n", topToBottom)
			}

			// Bottom to top
			if spaceRight && spaceUnder && spaceLeft && spaceUpper {
				bottomToTop = fmt.Sprintf("%s%s%s", string(lines[row+1][col-1:col]), string(char), string(lines[row-1][col+1:col+2]))
				fmt.Printf("Bottom to top: %s\n", bottomToTop)
			}

			if "MAS" == topToBottom && "MAS" == bottomToTop {
				answer2++
			}

			if "MAS" == topToBottom && "SAM" == bottomToTop {
				answer2++
			}

			if "SAM" == topToBottom && "SAM" == bottomToTop {
				answer2++
			}

			if "SAM" == topToBottom && "MAS" == bottomToTop {
				answer2++
			}
		}
	}

	fmt.Printf("Answer part 1: %d\n", answer1)
	fmt.Printf("Answer part 2: %d\n", answer2)
}
