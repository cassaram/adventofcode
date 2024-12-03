package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)

	fmt.Println("----- Part 1 -----")
	part1(input)

	fmt.Println("----- Part 2 -----")
	part2(input)
}

func part1(input string) {
	// Scan file for mul keywords
	scanArg1 := false
	scanArg2 := false
	buffer := make([]byte, 0, 32)
	sumOfMul := 0
	for i := 0; i < len(input); i++ {
		if len(buffer) == 0 {
			if input[i] == 'm' {
				buffer = append(buffer, input[i])
			}
			continue
		}
		if len(buffer) == 1 {
			if input[i] == 'u' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if len(buffer) == 2 {
			if input[i] == 'l' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if len(buffer) == 3 {
			if input[i] == '(' {
				buffer = append(buffer, input[i])
				scanArg1 = true
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if scanArg1 {
			if input[i] >= '0' && input[i] <= '9' {
				buffer = append(buffer, input[i])
			} else if input[i] == ',' {
				buffer = append(buffer, input[i])
				scanArg1 = false
				scanArg2 = true
			} else {
				buffer = make([]byte, 0, 32)
				scanArg1 = false
			}
			continue
		}
		if scanArg2 {
			if input[i] >= '0' && input[i] <= '9' {
				buffer = append(buffer, input[i])
			} else if input[i] == ')' {
				buffer = append(buffer, input[i])
				scanArg2 = false
				// Full mul command found, process it
				sumOfMul += processCommand(string(buffer))
				// Clear buffer
				buffer = make([]byte, 0, 32)
			} else {
				buffer = make([]byte, 0, 32)
				scanArg2 = false
			}
			continue
		}
	}
	fmt.Println("Sum of mul(a,b) commands:", sumOfMul)
}

func part2(input string) {
	// Scan file for mul(a,b) and do() and don't() keywords
	scanArg1 := false
	scanArg2 := false
	buffer := make([]byte, 0, 32)
	sumOfMul := 0
	mulEnabled := true
	for i := 0; i < len(input); i++ {
		if len(buffer) == 0 {
			if input[i] == 'm' && mulEnabled {
				buffer = append(buffer, input[i])
			}
			if input[i] == 'd' {
				buffer = append(buffer, input[i])
			}
			continue
		}
		// Handle mul(a,b) commands
		if string(buffer) == "m" {
			if input[i] == 'u' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "mu" {
			if input[i] == 'l' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "mul" {
			if input[i] == '(' {
				buffer = append(buffer, input[i])
				scanArg1 = true
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if scanArg1 {
			if input[i] >= '0' && input[i] <= '9' {
				buffer = append(buffer, input[i])
			} else if input[i] == ',' {
				buffer = append(buffer, input[i])
				scanArg1 = false
				scanArg2 = true
			} else {
				buffer = make([]byte, 0, 32)
				scanArg1 = false
			}
			continue
		}
		if scanArg2 {
			if input[i] >= '0' && input[i] <= '9' {
				buffer = append(buffer, input[i])
			} else if input[i] == ')' {
				buffer = append(buffer, input[i])
				scanArg2 = false
				// Full mul command found, process it
				sumOfMul += processCommand(string(buffer))
				// Clear buffer
				buffer = make([]byte, 0, 32)
			} else {
				buffer = make([]byte, 0, 32)
				scanArg2 = false
			}
			continue
		}
		// Handle do() and don't() commands
		if string(buffer) == "d" {
			if input[i] == 'o' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "do" {
			if input[i] == '(' {
				buffer = append(buffer, input[i])
			} else if input[i] == 'n' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "do(" {
			if input[i] == ')' {
				// Full do() command found, enable mul(a,b) commands
				mulEnabled = true
				// Reset buffer
				buffer = make([]byte, 0, 32)
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "don" {
			if input[i] == '\'' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "don'" {
			if input[i] == 't' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "don't" {
			if input[i] == '(' {
				buffer = append(buffer, input[i])
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
		if string(buffer) == "don't(" {
			if input[i] == ')' {
				// Full don't() command found, disable mul(a,b) commands
				mulEnabled = false
				// Reset buffer
				buffer = make([]byte, 0, 32)
			} else {
				buffer = make([]byte, 0, 32)
			}
			continue
		}
	}
	fmt.Println("Sum of mul(a,b) commands:", sumOfMul)
}

func processCommand(cmd string) int {
	// Trim command
	cmd = cmd[4:]
	cmd = cmd[:len(cmd)-1]
	// Split into arguments
	args := strings.Split(cmd, ",")
	arg1, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	arg2, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	return arg1 * arg2
}
