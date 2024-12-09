package main

import (
	"fmt"
	"os"
	"strings"
)

type direction uint8

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

func main() {
	guardMap, x, y := getInput()

	// Print unique cell count
	visitUniqueCount := doPathfinding(guardMap, x, y)
	fmt.Println("Unique positions visited:", visitUniqueCount)

	// Find all positions we can add an object such that it forces the guard into a loop
	obstructionCount := 0
	for i := range len(guardMap) {
		for j := range len(guardMap[0]) {
			// Check for existing obstruction
			if guardMap[i][j] {
				continue
			}
			// Check if guard started here
			if i == y && j == x {
				continue
			}
			// See if this causes an infinite loop
			testMap := copyMap(guardMap)
			testMap[i][j] = true
			if doPathfinding(testMap, x, y) == -1 {
				obstructionCount += 1
			}
		}
	}
	fmt.Println("Unique looping obstructions possible:", obstructionCount)
}

func copyMap(src [][]bool) [][]bool {
	dst := make([][]bool, len(src))
	for i := range src {
		dst[i] = make([]bool, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

// Return -1 if in infinite loop
func doPathfinding(guardMap [][]bool, x int, y int) int {
	visitMap := make(map[int]bool)
	visitDirMap := make(map[int]bool)
	visitUniqueCount := 0

	// Do pathfinding
	dir := UP
	for {
		mapIndex := (y << 8) | (x)
		mapDirIndex := (int(dir) << 16) | (y << 8) | (x)
		visitedDir, ok := visitDirMap[mapDirIndex]
		if ok && visitedDir {
			return -1
		}
		// Visit this location
		if !visitMap[int(mapIndex)] {
			visitUniqueCount += 1
			visitMap[int(mapIndex)] = true
			visitDirMap[mapDirIndex] = true
		}
		// Get next position
		nextX := -1
		nextY := -1
		switch dir {
		case UP:
			nextX = x
			nextY = y - 1
		case DOWN:
			nextX = x
			nextY = y + 1
		case LEFT:
			nextX = x - 1
			nextY = y
		case RIGHT:
			nextX = x + 1
			nextY = y
		}
		// Check if leaving map
		if nextX < 0 || nextX >= len(guardMap[0]) {
			break
		}
		if nextY < 0 || nextY >= len(guardMap) {
			break
		}
		// Check if next tile is not occupied
		if guardMap[nextY][nextX] {
			// Turn right
			switch dir {
			case UP:
				dir = RIGHT
			case DOWN:
				dir = LEFT
			case LEFT:
				dir = UP
			case RIGHT:
				dir = DOWN
			}
			continue
		}
		// Update next step
		x = nextX
		y = nextY
	}

	return visitUniqueCount
}

func getInput() ([][]bool, int, int) {
	// Get input file
	dataRaw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// Split into lines
	dataSliced := strings.Split(string(dataRaw), "\r\n")

	// Parse into map
	guardMap := make([][]bool, len(dataSliced))
	guardPosY := -1
	guardPosX := -1
	for i := range dataSliced {
		guardMap[i] = make([]bool, len(dataSliced[i]))
		for j := 0; j < len(dataSliced); j++ {
			if dataSliced[i][j] == '#' {
				guardMap[i][j] = true
				continue
			}
			if dataSliced[i][j] == '^' {
				guardPosY = i
				guardPosX = j
				continue
			}
		}
	}
	return guardMap, guardPosX, guardPosY
}
