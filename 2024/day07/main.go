package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result     int
	components []int
}

func main() {
	equations := getInput()

	totalCalibrationResult := 0
	for i := range equations {
		results := getResults(equations[i].components)
		for j := range results {
			if results[j] == equations[i].result {
				// Equation is valid
				totalCalibrationResult += equations[i].result
				break
			}
		}
	}
	fmt.Println("Total calibration result:", totalCalibrationResult)

	totalCalibrationResultConcat := 0
	for i := range equations {
		results := getResultsConcat(equations[i].components)
		for j := range results {
			if results[j] == equations[i].result {
				// Equation is valid
				totalCalibrationResultConcat += equations[i].result
				break
			}
		}
	}
	fmt.Println("Total calibration result concat:", totalCalibrationResultConcat)
}

func getResults(components []int) []int {
	if len(components) == 2 {
		return []int{components[0] + components[1], components[0] * components[1]}
	}
	passMul := []int{components[0] * components[1]}
	passMul = append(passMul, components[2:]...)
	passSum := []int{components[0] + components[1]}
	passSum = append(passSum, components[2:]...)
	resultsMul := getResults(passMul)
	resultsSum := getResults(passSum)
	return append(resultsSum, resultsMul...)
}

func getResultsConcat(components []int) []int {
	if len(components) == 2 {
		return []int{
			components[0] + components[1],
			components[0] * components[1],
			concat(components[0], components[1]),
		}
	}
	passMul := []int{components[0] * components[1]}
	passMul = append(passMul, components[2:]...)
	passSum := []int{components[0] + components[1]}
	passSum = append(passSum, components[2:]...)
	passCon := []int{concat(components[0], components[1])}
	passCon = append(passCon, components[2:]...)
	resultsMul := getResultsConcat(passMul)
	resultsSum := getResultsConcat(passSum)
	resultsCon := getResultsConcat(passCon)
	results := append(resultsSum, resultsMul...)
	return append(results, resultsCon...)
}

func getInput() []equation {
	// Get input file
	dataRaw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// Split into lines
	dataSliced := strings.Split(string(dataRaw), "\r\n")
	equations := make([]equation, len(dataSliced))
	for i := range dataSliced {
		mainParts := strings.Split(dataSliced[i], ": ")
		equations[i].result, err = strconv.Atoi(mainParts[0])
		if err != nil {
			panic(err)
		}
		componentsStr := strings.Split(mainParts[1], " ")
		equations[i].components = make([]int, len(componentsStr))
		for j := range componentsStr {
			equations[i].components[j], err = strconv.Atoi(componentsStr[j])
			if err != nil {
				panic(err)
			}
		}
	}
	return equations
}

func concat(a int, b int) int {
	result, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return result
}
