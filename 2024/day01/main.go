package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type number struct {
	number int
	index  int
}

func main() {
	input := getInput()

	// ----- Part 1 -----
	fmt.Println("----- Part 1 -----")

	// Parse lists
	leftList := make([]number, 0, 1024)
	rightList := make([]number, 0, 1024)
	for i, v := range input {
		vals := strings.Split(v, "   ")
		leftVal, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		rightVal, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, number{number: leftVal, index: i})
		rightList = append(rightList, number{number: rightVal, index: i})
	}

	// Sort values
	leftList = sortList(leftList)
	rightList = sortList(rightList)

	/*
		// Print values
		for i := range leftList {
			fmt.Println(leftList[i].number, rightList[i].number)
		}
	*/

	// Find distances and sum them
	sum := 0
	for i := range leftList {
		distance := (leftList[i].number - rightList[i].number)
		if distance < 0 {
			distance *= -1
		}
		sum += distance
	}

	fmt.Println("Sum of distances:", sum)

	// Part 2
	fmt.Println("----- Part 2 -----")

	// Get count of every number in right list
	countMap := make(map[int]int)
	for _, v := range rightList {
		countMap[v.number] += 1
	}

	// Get similarity score
	similarity := 0
	for _, v := range leftList {
		similarity += v.number * countMap[v.number]
	}

	fmt.Println("Similarity:", similarity)
}

func sortList(input []number) []number {
	sort.Slice(input, func(i, j int) bool {
		return input[i].number < input[j].number
	})
	return input
}

func getInput() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	dataSliced := strings.Split(string(data), "\r\n")
	return dataSliced
}
