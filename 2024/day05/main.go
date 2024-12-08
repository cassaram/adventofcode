package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Get input
	pages, rules := getInput()

	correctMiddleSum := 0
	correctedMiddleSum := 0
	for i := range pages {
		if isOrderedCorrect(pages[i], rules) {
			correctMiddleSum += getMiddleNum(pages[i])
			continue
		}
		corrected := correctOrder(pages[i], rules)
		correctedMiddleSum += getMiddleNum(corrected)
	}
	fmt.Println("Sum of correct middle pages:  ", correctMiddleSum)
	fmt.Println("Sum of corrected middle pages:", correctedMiddleSum)
}

func correctOrder(page []int, rules map[int][]int) []int {
	// First element is always correct
	for i := 1; i < len(page); i++ {
		for j := range rules[page[i]] {
			for k := 0; k < i; k++ {
				if page[k] == rules[page[i]][j] {
					return correctOrder(reorderSingle(page, i, k), rules)
				}
			}
		}
	}
	return page
}

func reorderSingle(page []int, srcIdx int, dstIdx int) []int {
	// Copy element to insert
	ins := page[srcIdx]
	// Copy elements at insert to i+1
	for i := srcIdx - 1; i >= dstIdx; i-- {
		page[i+1] = page[i]
	}
	page[dstIdx] = ins
	return page
}

func isOrderedCorrect(page []int, rules map[int][]int) bool {
	// First element is always correct
	for i := 1; i < len(page); i++ {
		for j := range rules[page[i]] {
			for k := 0; k < i; k++ {
				if page[k] == rules[page[i]][j] {
					return false
				}
			}
		}
	}
	return true
}

func getMiddleNum(page []int) int {
	return page[len(page)/2]
}

func getInput() ([][]int, map[int][]int) {
	// Get input file
	dataRaw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// Split into lines
	dataSliced := strings.Split(string(dataRaw), "\r\n")

	// Read ordering rules
	rules := make(map[int][]int)
	pagesStart := -1
	for i, line := range dataSliced {
		// Check if this line is the separation
		if len(line) == 0 {
			pagesStart = i + 1
			break
		}

		// Get both numbers
		nums := strings.Split(line, "|")
		key, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		val, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		if len(rules[key]) == 0 {
			rules[key] = make([]int, 0, 24)
		}
		rules[key] = append(rules[key], val)
	}

	// Read pages
	pages := make([][]int, len(dataSliced)-pagesStart)
	for i := 0; i < len(pages); i++ {
		line := dataSliced[i+pagesStart]
		nums := strings.Split(line, ",")
		pages[i] = make([]int, len(nums))
		for j := range len(nums) {
			pages[i][j], err = strconv.Atoi(nums[j])
			if err != nil {
				panic(err)
			}
		}
	}
	return pages, rules
}
