// [Advent of Code: 2024: Day 2: Red-Nosed Reports]
//
// Part 1 Example Expected: 2
// Part 2 Example Expected: 4
//
// [Advent of Code: 2024: Day 2: Red-Nosed Reports]: https://adventofcode.com/2024/day/2

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	safeReports := 0

	reports := processInput()

	for _, report := range reports {
		if isReportValid(report) {
			safeReports++
		}
	}

	return safeReports
}

func part2() int {
	safeReports := 0

	reports := processInput()

	for _, report := range reports {
		if isReportValidWithTolerance(report) {
			safeReports++
		}
	}

	return safeReports
}

func processInput() [][]int {
	file, err := os.ReadFile("./2024/02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	formatInput := strings.Split(strings.TrimSpace(string(file)), "\n")
	var reports [][]int
	for _, v := range formatInput {
		var report []int
		_ = json.Unmarshal([]byte("["+strings.ReplaceAll(v, " ", ",")+"]"), &report)
		reports = append(reports, report)
	}

	return reports
}

func isReportValid(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		validReportLevels := validateReportLevels(report)

		if !validReportLevels {
			return false
		}
	}
	return true
}

func validateReportLevels(report []int) bool {
	var decreasingDirection bool
	for i := 1; i < len(report); i++ {
		if i == 1 {
			decreasingDirection = report[i-1] > report[i]
		}
		stillDecreasing := report[i-1] > report[i]
		if decreasingDirection != stillDecreasing {
			return false
		}
		difference := absInt(report[i-1], report[i])
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func absInt(a, b int) int {
	return maxInt(a, b) - minInt(a, b)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isReportValidWithTolerance(report []int) bool {
	numInvalidLevels, invalidLevelLocations := invalidReportStats(report)

	if numInvalidLevels == 0 {
		return true
	}
	for _, v := range invalidLevelLocations {

		filteredSlice := removeIndex(report, v)

		if isReportValid(filteredSlice) {
			return true
		}
	}
	return false
}

func invalidReportStats(report []int) (numInvalidLevels int, invalidLevelLocations []int) {
	totalInvalidLevels := 0

	for i := 1; i < len(report); i++ {
		validReportLevels := validateReportLevels(report)
		if !validReportLevels {
			invalidLevelLocations = append(invalidLevelLocations, i-1, i)
			totalInvalidLevels++
			continue
		}
	}
	return totalInvalidLevels, slices.Clip(invalidLevelLocations)
}

func removeIndex(slice []int, index int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	newSlice = slices.Clip(slices.Delete(newSlice, index, index+1))
	return newSlice
}
