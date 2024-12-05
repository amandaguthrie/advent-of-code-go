// [Advent of Code: 2024: Day 6: _]
//
// Part 1 Example Expected:
// Part 2 Example Expected:
//
// [Advent of Code: 2024: Day 6: _]: https://adventofcode.com/2024/day/6

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %s\n", part1())
	fmt.Printf("Part 2: %s\n", part2())
}

func part1() []string {

	input := processInput()

	return input
}

func part2() []string {
	input := processInput()

	return input
}

func processInput() []string {
	file, err := os.ReadFile("./2024/06/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	formatInput := strings.Split(string(file), "\n")
	return formatInput
}
