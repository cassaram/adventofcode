package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()

	// ----- Part 1 -----
	fmt.Println("----- Part 1 -----")

	safeCount := 0
	for i := range input {
		if determineIfSafe(parseLine(input[i])) {
			safeCount += 1
		}
	}

	fmt.Println("Number of safe reports:", safeCount)

	// ----- Part 2 -----
	fmt.Println("----- Part 2 -----")
	safeCount2 := 0
	for i := range input {
		if doDampener(parseLine(input[i])) {
			safeCount2 += 1
		}
	}

	fmt.Println("Number of safe reports:", safeCount2)

}

func getInput() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	dataSliced := strings.Split(string(data), "\r\n")
	return dataSliced
}

func parseLine(line string) []int {
	vals := strings.Split(line, " ")
	result := make([]int, 0, 16)
	for _, val := range vals {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		result = append(result, valInt)
	}
	return result
}

func determineIfSafe(line []int) bool {
	// Ensure first two values are different
	if abs(line[0]-line[1]) < 1 || abs(line[0]-line[1]) > 3 {
		return false
	}

	// Determine if increasing or decreasing
	inc := true
	if line[0] > line[1] {
		inc = false
	}

	// Step through all values
	for i := 2; i < len(line); i++ {
		if abs(line[i-1]-line[i]) < 1 || abs(line[i-1]-line[i]) > 3 {
			return false
		}
		if inc {
			if line[i-1]-line[i] > 0 {
				return false
			}
		} else {
			if line[i-1]-line[i] < 0 {
				return false
			}
		}
	}
	return true
}

func doDampener(line []int) bool {
	// Brute force this because it is the easiest way to do it
	// Computational expense should be only midly worse compared to doing it in one pass
	// Allows for processor optimizations as it does not delay the pipeline thanks to branch prediction
	if determineIfSafe(line) {
		return true
	}
	for i := range line {
		testLine := make([]int, 0, len(line)-1)
		for j := range line {
			if j != i {
				testLine = append(testLine, line[j])
			}
		}
		if determineIfSafe(testLine) {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
