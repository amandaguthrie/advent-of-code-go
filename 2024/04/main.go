// [Advent of Code: 2024: Day 4: Ceres Search]
//
// Part 1 Example Expected: 18
// Part 2 Example Expected: 9
//
// [Advent of Code: 2024: Day 4: Ceres Search]: https://adventofcode.com/2024/day/4

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var reXmas = regexp.MustCompile("(XMAS)")
var reXmasReverse = regexp.MustCompile("(SAMX)")

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	matches := 0
	crossword := processInput()

	matches += processDirections(crossword)

	return matches
}

func part2() int {
	matches := 0
	crossword := processInput()

	matches += findMasX(crossword)

	return matches
}

func processInput() [][]string {
	file, err := os.ReadFile("./2024/04/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	formatInput := strings.Split(string(file), "\r\n")
	parsedInput := make([][]string, len(formatInput))
	for i, v := range formatInput {
		parsedInput[i] = strings.Split(v, "")
	}
	return parsedInput
}

func matchForwardAndReverse(s string) int {
	matches := 0
	matches += len(reXmas.FindAllString(s, -1))
	matches += len(reXmasReverse.FindAllString(s, -1))
	return matches
}

func processDirections(crossword [][]string) int {
	matches := 0
	for rowIndex := 0; rowIndex < len(crossword); rowIndex++ {
		row := ""
		if rowIndex == 0 {
			for colIndex := range len(crossword[rowIndex]) {
				column := ""
				diagonal := ""
				fromEndDiagonal := ""
				diagonalIndex := 0
				row += crossword[rowIndex][colIndex]
				for colRowIndex := range len(crossword) {
					column += crossword[colRowIndex][colIndex]
					if colIndex+diagonalIndex < len(crossword[colRowIndex]) {
						diagonal += crossword[colRowIndex][colIndex+diagonalIndex]
						fromEndDiagonal += crossword[colRowIndex][len(crossword[colRowIndex])-1-colIndex-diagonalIndex]
					}

					diagonalIndex++
				}
				matches += matchForwardAndReverse(row)
				matches += matchForwardAndReverse(column)
				matches += matchForwardAndReverse(diagonal)
				matches += matchForwardAndReverse(fromEndDiagonal)
			}
		} else {
			diagonal := ""
			fromEndDiagonal := ""
			for colIndex := 0; colIndex < len(crossword); colIndex++ {
				row += crossword[rowIndex][colIndex]
				if colIndex < len(crossword)-rowIndex {
					diagonal += crossword[rowIndex+colIndex][colIndex]
					fromEndDiagonal += crossword[rowIndex+colIndex][len(crossword[rowIndex])-1-colIndex]
				}
			}
			matches += matchForwardAndReverse(row)
			matches += matchForwardAndReverse(diagonal)
			matches += matchForwardAndReverse(fromEndDiagonal)
		}
	}
	return matches

}

func findMasX(crossword [][]string) int {
	matches := 0
	for rowIndex, row := range crossword {
		for colIndex, cell := range row {
			if cell != "A" {
				continue
			}
			leftHook := ""
			rightHook := ""
			if rowIndex-1 >= 0 && colIndex-1 >= 0 && rowIndex+1 < len(crossword) && colIndex+1 < len(crossword[rowIndex]) {
				leftHook += crossword[rowIndex-1][colIndex-1] + crossword[rowIndex+1][colIndex+1]
				if leftHook != "SM" && leftHook != "MS" {
					continue
				}
				rightHook += crossword[rowIndex-1][colIndex+1] + crossword[rowIndex+1][colIndex-1]
				if rightHook != "SM" && rightHook != "MS" {
					continue
				}
				matches++
			}

		}
	}
	return matches
}
