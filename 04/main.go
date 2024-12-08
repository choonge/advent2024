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
			if char != 'X' {
				continue
			}

			fmt.Printf("Row %d | Col %d | %c\n", row, col, char)

			spaceRight := col+3 < len(line)
			spaceUnder := row+3 < len(lines)
			spaceLeft := col-3 >= 0
			spaceUpper := row-3 >= 0

			// Right
			if spaceRight {
				right := string(line[col : col+4])
				fmt.Printf("Right: %s\n", right)

				if "XMAS" == right {
					answer1++
				}
			}

			// Right-Under
			if spaceRight && spaceUnder {
				rightUnder := fmt.Sprintf("%s%s%s%s", string(lines[row][col:col+1]), string(lines[row+1][col+1:col+2]), string(lines[row+2][col+2:col+3]), string(lines[row+3][col+3:col+4]))
				fmt.Printf("Right-Under: %s\n", rightUnder)

				if "XMAS" == rightUnder {
					answer1++
				}
			}

			// Under
			if spaceUnder {
				under := fmt.Sprintf("%s%s%s%s", string(lines[row][col:col+1]), string(lines[row+1][col:col+1]), string(lines[row+2][col:col+1]), string(lines[row+3][col:col+1]))
				fmt.Printf("Under: %s\n", under)

				if "XMAS" == under {
					answer1++
				}
			}

			// Left-Under
			if spaceLeft && spaceUnder {
				leftUnder := fmt.Sprintf("%s%s%s%s", string(lines[row][col:col+1]), string(lines[row+1][col-1:col]), string(lines[row+2][col-2:col-1]), string(lines[row+3][col-3:col-2]))
				fmt.Printf("Left-Under: %s\n", leftUnder)

				if "XMAS" == leftUnder {
					answer1++
				}
			}

			// Left
			if spaceLeft {
				left := fmt.Sprintf("%s%s%s%s", string(lines[row][col:col+1]), string(lines[row][col-1:col]), string(lines[row][col-2:col-1]), string(lines[row][col-3:col-2]))
				fmt.Printf("Left: %s\n", left)

				if "XMAS" == left {
					answer1++
				}
			}

			// Left-Upper
			if spaceLeft && spaceUpper {
				leftUpper := fmt.Sprintf("%s%s%s%s", string(lines[row][col:col+1]), string(lines[row-1][col-1:col]), string(lines[row-2][col-2:col-1]), string(lines[row-3][col-3:col-2]))
				fmt.Printf("Left-Upper: %s\n", leftUpper)

				if "XMAS" == leftUpper {
					answer1++
				}
			}

			// Upper
			if spaceUpper {
				upper := fmt.Sprintf("%s%s%s%s", string(lines[row][col:col+1]), string(lines[row-1][col:col+1]), string(lines[row-2][col:col+1]), string(lines[row-3][col:col+1]))
				fmt.Printf("Upper: %s\n", upper)

				if "XMAS" == upper {
					answer1++
				}
			}

			// Right-Upper
			if spaceUpper && spaceRight {
				rightUpper := fmt.Sprintf("%s%s%s%s", string(lines[row][col:col+1]), string(lines[row-1][col+1:col+2]), string(lines[row-2][col+2:col+3]), string(lines[row-3][col+3:col+4]))
				fmt.Printf("Right-Upper: %s\n", rightUpper)

				if "XMAS" == rightUpper {
					answer1++
				}
			}
		}
	}

	fmt.Printf("Answer part 1: %d\n", answer1)
	fmt.Printf("Answer part 2: %d\n", answer2)
}
