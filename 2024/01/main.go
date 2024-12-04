// [Advent of Code: 2024: Day 1: Historian Hysteria]
//
// Part 1 Example Expected: 11
// Part 2 Example Expected: 31
//
// [Advent of Code: 2024: Day 1: Historian Hysteria]: https://adventofcode.com/2024/day/1

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func processInput() ([]int, []int) {
	file, err := os.ReadFile("./2024/01/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	formatInput := strings.Split(string(file), "\n")
	list1 := make([]int, len(formatInput))
	list2 := make([]int, len(formatInput))

	for i, v := range formatInput {
		list1[i], _ = strconv.Atoi(strings.Fields(v)[0])
		list2[i], _ = strconv.Atoi(strings.Fields(v)[1])
	}

	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2
}

// Pair up the smallest number in the left list with the smallest number in the right list, and so on.
//
// Within each pair, figure out how apart the two numbers are and add up all the distances.
func part1() int {
	totalDistances := 0

	list1, list2 := processInput()

	for i, v := range list1 {
		difference := 0
		if v < list2[i] {
			difference = list2[i] - v
		} else {
			difference = v - list2[i]
		}
		totalDistances += difference
	}

	return totalDistances
}

func part2() int {
	similarityScore := 0

	list1, list2 := processInput()

	list2Occurrences := make(map[int]int)

	// Gather counts for each value in list 2
	for _, v := range list2 {
		list2Occurrences[v]++
	}

	// Add value * occurrences to similarity score
	for _, v := range list1 {
		similarityScore += v * list2Occurrences[v]
	}

	return similarityScore
}
