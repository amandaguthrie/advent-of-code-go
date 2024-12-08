// [Advent of Code: 2024: Day 7: Bridge Repair]
//
// Part 1 Example Expected: 3749
// Part 2 Example Expected: 11387
//
// [Advent of Code: 2024: Day 7: Bridge Repair]: https://adventofcode.com/2024/day/7

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {

	calibrationEquations := processInput()
	total := 0
	for _, equation := range calibrationEquations {
		isValid, _ := isValidCalibration(equation, false)
		if isValid {
			total += equation.total
		}
	}

	return total
}

func part2() int {
	calibrationEquations := processInput()
	total := 0
	for _, equation := range calibrationEquations {
		isValid, _ := isValidCalibration(equation, true)
		if isValid {
			total += equation.total
		}
	}

	return total
}

type Equation struct {
	numbers []int
	total   int
	valid   *bool
}

func processInput() (calibrationEquations []Equation) {
	file, err := os.ReadFile("./2024/07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(file), "\r\n")
	for _, row := range rows {
		totalNumSplit := strings.Split(row, ":")
		total, _ := strconv.Atoi(totalNumSplit[0])
		numStrings := strings.Fields(totalNumSplit[1])
		nums := make([]int, len(numStrings))
		for s, numString := range numStrings {
			nums[s], _ = strconv.Atoi(numString)
		}
		calibrationEquations = append(calibrationEquations, Equation{nums, total, nil})
	}
	return calibrationEquations
}

func isValidCalibration(equation Equation, part2 bool) (bool, Equation) {
	sumResult := calculateResult(equation, 0, "+", part2)
	productResult := calculateResult(equation, 0, "*", part2)
	concatResult := false
	if part2 {
		concatResult = calculateResult(equation, 0, "||", part2)
	}

	result := sumResult || productResult || concatResult
	equation.valid = &result
	return result, equation
}

func calculateResult(equation Equation, carryTotal int, operation string, part2 bool) bool {
	var totalResult int
	if operation == "+" {
		totalResult = carryTotal + equation.numbers[0]
	} else if operation == "*" {
		totalResult = carryTotal * equation.numbers[0]
	} else if operation == "||" {
		newNumAsString := strconv.Itoa(carryTotal) + strconv.Itoa(equation.numbers[0])
		concatResult, _ := strconv.Atoi(newNumAsString)
		totalResult = concatResult
	}

	if len(equation.numbers) == 1 {
		return totalResult == equation.total
	}

	sumResult := calculateResult(Equation{equation.numbers[1:], equation.total, equation.valid}, totalResult, "+", part2)
	productResult := calculateResult(Equation{equation.numbers[1:], equation.total, equation.valid}, totalResult, "*", part2)
	concatResult := false
	if part2 {
		concatResult = calculateResult(Equation{equation.numbers[1:], equation.total, equation.valid}, totalResult, "||", part2)
	}

	return sumResult || productResult || concatResult
}
