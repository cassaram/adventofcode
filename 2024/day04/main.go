package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get input
	input := getInput()

	part1(input)
	part2(input)
}

func part1(input [][]byte) {
	fmt.Println("---- Part 1 ----")
	// Patterns to search for
	patterns := [][][]byte{
		{
			{'X', 'M', 'A', 'S'},
		},
		{
			{'S', 'A', 'M', 'X'},
		},
		{
			{'X'},
			{'M'},
			{'A'},
			{'S'},
		},
		{
			{'S'},
			{'A'},
			{'M'},
			{'X'},
		},
		{
			{'X', '.', '.', '.'},
			{'.', 'M', '.', '.'},
			{'.', '.', 'A', '.'},
			{'.', '.', '.', 'S'},
		},
		{
			{'S', '.', '.', '.'},
			{'.', 'A', '.', '.'},
			{'.', '.', 'M', '.'},
			{'.', '.', '.', 'X'},
		},
		{
			{'.', '.', '.', 'S'},
			{'.', '.', 'A', '.'},
			{'.', 'M', '.', '.'},
			{'X', '.', '.', '.'},
		},
		{
			{'.', '.', '.', 'X'},
			{'.', '.', 'M', '.'},
			{'.', 'A', '.', '.'},
			{'S', '.', '.', '.'},
		},
	}

	count := 0
	for i := range patterns {
		count += searchWithPattern(input, patterns[i])
	}
	fmt.Println("Count of XMAS variants:", count)
}

func part2(input [][]byte) {
	fmt.Println("---- Part 2 ----")
	// Patterns to search for
	patterns := [][][]byte{
		{
			{'M', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'S'},
		},
		{
			{'S', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'M'},
		},
		{
			{'M', '.', 'S'},
			{'.', 'A', '.'},
			{'M', '.', 'S'},
		},
		{
			{'S', '.', 'M'},
			{'.', 'A', '.'},
			{'S', '.', 'M'},
		},
	}

	count := 0
	for i := range patterns {
		count += searchWithPattern(input, patterns[i])
	}
	fmt.Println("Count of X-MAS variants:", count)
}

func searchWithPattern(input [][]byte, pattern [][]byte) int {
	// Variable to store how many instances we find
	count := 0
	for i := 0; i < len(input)-len(pattern)+1; i++ {
		for j := 0; j < len(input[0])-len(pattern[0])+1; j++ {
			// Get subset area
			subset := make([][]byte, len(pattern))
			for k := range subset {
				subset[k] = input[i+k][j : j+len(pattern[0])]
			}
			if isSimilar(subset, pattern) {
				count += 1
			}
		}
	}

	return count
}

func isSimilar(a [][]byte, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[0]); j++ {
			if b[i][j] == '.' {
				continue
			}
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func getInput() [][]byte {
	// Get input file
	dataRaw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// Split into lines
	dataSliced := strings.Split(string(dataRaw), "\r\n")
	// Create 2d array
	data := make([][]byte, len(dataSliced))
	for i := range dataSliced {
		data[i] = []byte(dataSliced[i])
	}
	return data
}
