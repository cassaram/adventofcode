package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get input
	input := getInput()

	// Get map of transmitters on each frequency
	antennas := getAntennaMap(input)

	part1(input, antennas)
	part2(input, antennas)
}

func part1(input [][]rune, antennas map[rune][][]int) {
	// Get all antinodes
	antinodes := make([][]bool, len(input))
	for i := range input {
		antinodes[i] = make([]bool, len(input[i]))
	}
	antinodesCount := 0
	for key := range antennas {
		for i := range antennas[key] {
			locA := antennas[key][i]
			for j := i + 1; j < len(antennas[key]); j++ {
				locB := antennas[key][j]
				nodes := getAntinodes(locA, locB)
				loc1 := nodes[0]
				loc2 := nodes[1]
				// Ensure indexes are in bounds
				if loc1[0] >= 0 && loc1[0] < len(antinodes) && loc1[1] >= 0 && loc1[1] < len(antinodes[1]) {
					if !antinodes[loc1[0]][loc1[1]] {
						antinodes[loc1[0]][loc1[1]] = true
						antinodesCount += 1
					}
				}
				if loc2[0] >= 0 && loc2[0] < len(input) && loc2[1] >= 0 && loc2[1] < len(input[1]) {
					if !antinodes[loc2[0]][loc2[1]] {
						antinodes[loc2[0]][loc2[1]] = true
						antinodesCount += 1
					}
				}
			}
		}
	}
	fmt.Println("Number of unique antinodes:", antinodesCount)
}

func part2(input [][]rune, antennas map[rune][][]int) {
	// Get all antinodes
	antinodes := make([][]bool, len(input))
	for i := range input {
		antinodes[i] = make([]bool, len(input[i]))
	}
	antinodesCount := 0
	for key := range antennas {
		for i := range antennas[key] {
			locA := antennas[key][i]
			for j := i + 1; j < len(antennas[key]); j++ {
				locB := antennas[key][j]
				nodes := getAntinodes2(locA, locB, len(input), len(input[0]))
				for i := range nodes {
					if nodes[i][0] >= 0 && nodes[i][0] < len(antinodes) && nodes[i][1] >= 0 && nodes[i][1] < len(antinodes[1]) {
						if !antinodes[nodes[i][0]][nodes[i][1]] {
							antinodes[nodes[i][0]][nodes[i][1]] = true
							antinodesCount += 1
						}
					}
				}
			}
		}
	}
	fmt.Println("Number of unique antinodes:", antinodesCount)
}

func getAntinodes(a []int, b []int) [][]int {
	result := make([][]int, 2)
	result[0] = make([]int, 2)
	result[1] = make([]int, 2)
	di := abs(a[0] - b[0])
	dj := abs(a[1] - b[1])
	// Switch depending on quadrant
	if (a[0]-b[0]) > 0 && (a[1]-b[1]) > 0 {
		// Down, Right
		result[0] = []int{b[0] - di, b[1] - dj}
		result[1] = []int{a[0] + di, a[1] + dj}
		return result
	}
	if (a[0]-b[0]) > 0 && (a[1]-b[1]) <= 0 {
		// Down, Left
		result[0] = []int{b[0] - di, b[1] + dj}
		result[1] = []int{a[0] + di, a[1] - dj}
		return result
	}
	if (a[0]-b[0]) <= 0 && (a[1]-b[1]) > 0 {
		// Up, Right
		result[0] = []int{a[0] - di, a[1] + dj}
		result[1] = []int{b[0] + di, b[1] - dj}
		return result
	}
	if (a[0]-b[0]) <= 0 && (a[1]-b[1]) <= 0 {
		// Up, Left
		result[0] = []int{a[0] - di, a[1] - dj}
		result[1] = []int{b[0] + di, b[1] + dj}
		return result
	}
	return result
}

func isValidLocation(loc []int, iMax int, jMax int) bool {
	if loc[0] < 0 {
		return false
	}
	if loc[0] >= iMax {
		return false
	}
	if loc[1] < 0 {
		return false
	}
	if loc[1] >= jMax {
		return false
	}
	return true
}

func getAntinodes2(a []int, b []int, iMax int, jMax int) [][]int {
	result := make([][]int, 0)
	di := abs(a[0] - b[0])
	dj := abs(a[1] - b[1])
	// Switch depending on quadrant
	if (a[0]-b[0]) > 0 && (a[1]-b[1]) > 0 {
		// Down, Right
		i := b[0]
		j := b[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i -= di
			j -= dj
		}
		i = a[0]
		j = a[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i += di
			j += dj
		}
		return result
	}
	if (a[0]-b[0]) > 0 && (a[1]-b[1]) <= 0 {
		// Down, Left
		i := b[0]
		j := b[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i -= di
			j += dj
		}
		i = a[0]
		j = a[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i += di
			j -= dj
		}
		return result
	}
	if (a[0]-b[0]) <= 0 && (a[1]-b[1]) > 0 {
		// Up, Right
		i := a[0]
		j := a[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i -= di
			j += dj
		}
		i = b[0]
		j = b[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i += di
			j -= dj
		}
		return result
	}
	if (a[0]-b[0]) <= 0 && (a[1]-b[1]) <= 0 {
		// Up, Left
		i := a[0]
		j := a[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i -= di
			j -= dj
		}
		i = b[0]
		j = b[1]
		for {
			loc := []int{i, j}
			if !isValidLocation(loc, iMax, jMax) {
				break
			}
			result = append(result, loc)
			i += di
			j += dj
		}
		return result
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func getAntennaMap(input [][]rune) map[rune][][]int {
	antMap := make(map[rune][][]int)
	for i := range input {
		for j := range input[i] {
			if input[i][j] != '.' {
				antMap[input[i][j]] = append(antMap[input[i][j]], []int{i, j})
			}
		}
	}
	return antMap
}

func getInput() [][]rune {
	// Get input file
	dataRaw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// Split into lines
	dataSliced := strings.Split(string(dataRaw), "\r\n")
	result := make([][]rune, len(dataSliced))
	for i := range dataSliced {
		result[i] = []rune(dataSliced[i])
	}
	return result
}
