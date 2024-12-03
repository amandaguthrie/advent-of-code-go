// [Advent of Code: 2024: Day 3: Mull It Over]
//
// Part 1 Example Expected: 161
// Part 2 Example Expected: 48
//
// [Advent of Code: 2024: Day 3: Mull It Over]: https://adventofcode.com/2024/day/3

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	muls := processInput("./2024/03/example-part1.txt", false)

	return sumProducts(muls)
}

func part2() int {
	muls := processInput("./2024/03/example-part2.txt", true)

	return sumProducts(muls)
}

func processInput(filePath string, part2 bool) [][]string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	parsed := string(file)
	if part2 {
		reRemoveDont := regexp.MustCompile(`(?s)don't\(\).*?do\(\)|$`)
		parsed = reRemoveDont.ReplaceAllString(parsed, "")
	}

	reMul := regexp.MustCompile(`mul\((?P<mul1>\d{1,3}),(?P<mul2>\d{1,3})\)`)
	return reMul.FindAllStringSubmatch(parsed, -1)
}

func sumProducts(muls [][]string) int {
	mulTotal := 0

	for _, mul := range muls {
		num1, _ := strconv.Atoi(mul[1])
		num2, _ := strconv.Atoi(mul[2])
		mulTotal += num1 * num2
	}

	return mulTotal
}
