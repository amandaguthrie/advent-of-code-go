// [Advent of Code: 2024: Day 5: Print Queue]
//
// Part 1 Example Expected: 143
// Part 2 Example Expected: 123
//
// [Advent of Code: 2024: Day 5: Print Queue]: https://adventofcode.com/2024/day/5

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

	rules, updatePages := processInput()

	var correctlyOrderedUpdates [][]int

	for _, update := range updatePages {
		validUpdate := isValidUpdate(rules, update)
		if validUpdate {
			correctlyOrderedUpdates = append(correctlyOrderedUpdates, update)
		}
	}

	return sumMiddlePageNumbers(correctlyOrderedUpdates)
}

func part2() int {
	rules, updatePages := processInput()

	var incorrectlyOrderedUpdates [][]int

	for _, update := range updatePages {
		validUpdate := isValidUpdate(rules, update)
		if !validUpdate {
			incorrectlyOrderedUpdates = append(incorrectlyOrderedUpdates, update)
		}
	}

	for _, update := range incorrectlyOrderedUpdates {
		slices.SortFunc(update, func(a, b int) int {
			prereqs := rules[a]
			if slices.Contains(prereqs, b) {
				return 1
			}
			return -1
		})
	}

	return sumMiddlePageNumbers(incorrectlyOrderedUpdates)
}

func processInput() (rules map[int][]int, updatePages [][]int) {
	file, err := os.ReadFile("./2024/05/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	separateInput := strings.Split(string(file), "\r\n\r\n")
	rulesInput := strings.Split(separateInput[0], "\r\n")
	rules = make(map[int][]int)
	for _, rule := range rulesInput {
		var pageToPrint int
		var whenBeforePage int
		_, _ = fmt.Sscanf(rule, "%d|%d", &pageToPrint, &whenBeforePage)
		rules[pageToPrint] = append(rules[pageToPrint], whenBeforePage)
	}

	updateInput := strings.Split(separateInput[1], "\r\n")
	for _, update := range updateInput {
		var updateResult []int
		_ = json.Unmarshal([]byte("["+update+"]"), &updateResult)
		updatePages = append(updatePages, updateResult)
	}

	return rules, updatePages
}

func filterIrrelevantRules(baseMap map[int][]int, update []int) (filteredMap map[int][]int) {
	filteredMap = make(map[int][]int)
	for page, preReqs := range baseMap {
		for _, preReq := range preReqs {
			if slices.Contains(update, preReq) {
				filteredMap[page] = append(filteredMap[page], preReq)
			}
		}
	}
	return filteredMap
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	slices.Reverse(update)
	filterRules := filterIrrelevantRules(rules, update)

	for pageIndex, page := range update {
		restOfPages := update[pageIndex+1:]
		for _, p := range restOfPages {
			if slices.Contains(filterRules[page], p) {
				return false
			}
		}
	}
	return true
}

func sumMiddlePageNumbers(updates [][]int) (sum int) {
	for _, update := range updates {
		// All update lengths are odd, otherwise would've needed an odd/even check here.
		sum += update[len(update)/2]
	}
	return sum
}
