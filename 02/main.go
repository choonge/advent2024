package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	answer1 := 1000
	answer2 := 1000

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		report := strings.Fields(scanner.Text())

		fmt.Printf("First try: %v\n", report)
		isUnsafe, unsafePosition := getIsUnsafe(report)
		if isUnsafe == false {
			fmt.Printf("Safe\n\n")
			continue
		}

		answer1--

		// Try without the unsafe position
		reportUnsafeCurrent := make([]string, len(report))
		copy(reportUnsafeCurrent, report)
		reportUnsafeCurrent = slices.Delete(reportUnsafeCurrent, unsafePosition, unsafePosition+1)
		fmt.Printf("Second try: %v\n", reportUnsafeCurrent)
		isUnsafe, _ = getIsUnsafe(reportUnsafeCurrent)
		if isUnsafe == false {
			fmt.Printf("Safe\n\n")
			continue
		}

		// Try without the unsafe position + 1
		reportUnsafeNext := make([]string, len(report))
		copy(reportUnsafeNext, report)
		reportUnsafeNext = slices.Delete(reportUnsafeNext, unsafePosition+1, unsafePosition+2)
		fmt.Printf("Second try: %v\n", reportUnsafeNext)
		isUnsafe, _ = getIsUnsafe(reportUnsafeNext)
		if isUnsafe == false {
			fmt.Printf("Safe\n\n")
			continue
		}

		// Try without the first position
		reportUnsafeFirst := make([]string, len(report))
		copy(reportUnsafeFirst, report)
		reportUnsafeFirst = slices.Delete(reportUnsafeFirst, 0, 1)
		fmt.Printf("Third try: %v\n", reportUnsafeFirst)
		isUnsafe, _ = getIsUnsafe(reportUnsafeFirst)
		if isUnsafe == false {
			fmt.Printf("Safe\n\n")
			continue
		}

		fmt.Printf("Unsafe\n\n")
		answer2--
	}

	// Part one
	fmt.Printf("Answer part 1: %d\n", answer1)

	// Part two
	fmt.Printf("Answer part 2: %d\n", answer2)
}

func getIsUnsafe(fields []string) (bool, int) {
	order := ""

	for k, v := range fields {
		current, _ := strconv.Atoi(v)

		if k >= len(fields)-1 {
			break
		}

		next, _ := strconv.Atoi(fields[k+1])
		diff := next - current

		// The levels are either all increasing or all decreasing.
		if order == "" {
			if diff > 0 {
				order = "increasing"
			} else {
				order = "decreasing"
			}
		} else {
			if order == "increasing" && diff < 0 {
				return true, k
			}

			if order == "decreasing" && diff > 0 {
				return true, k
			}
		}

		if diff < 0 {
			diff = -diff
		}

		// Any two adjacent levels differ by at least one and at most three.
		if diff < 1 || diff > 3 {
			return true, k
		}
	}

	return false, 0
}
