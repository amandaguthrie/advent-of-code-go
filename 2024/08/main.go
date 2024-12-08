// [Advent of Code: 2024: Day 8: Resonant Collinearity]
//
// Part 1 Example Expected: 14
// Part 2 Example Expected: 34
//
// [Advent of Code: 2024: Day 8: Resonant Collinearity]: https://adventofcode.com/2024/day/8

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type MapPosition struct {
	row int
	col int
}

type GridPosition struct {
	row       int
	col       int
	antenna   *string
	antinodes []string
}

type AntennaMap map[string][]MapPosition

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {

	frequencyMap, grid, _ := processInput()
	grid = addAntinodes(frequencyMap, grid, false)
	antinodeCount := countAntinodes(grid)

	return antinodeCount
}

func part2() int {
	frequencyMap, grid, _ := processInput()
	grid = addAntinodes(frequencyMap, grid, true)
	antinodeCount := countAntinodes(grid)

	return antinodeCount
}

func processInput() (antennaMap AntennaMap, grid []GridPosition, maxGridSize MapPosition) {
	antennaMap = make(AntennaMap)
	grid = make([]GridPosition, 0)
	file, err := os.ReadFile("./2024/08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	formatInput := strings.Split(string(file), "\r\n")
	maxGridSize.row = len(formatInput)
	for rowIndex, row := range formatInput {
		rowChars := strings.Split(row, "")
		maxGridSize.col = len(rowChars)
		for colIndex, char := range rowChars {
			if char != "." {
				if antennaMap[char] == nil {
					antennaMap[char] = make([]MapPosition, 0)
				}
				antennaMap[char] = append(antennaMap[char], MapPosition{rowIndex, colIndex})
				grid = append(grid, GridPosition{rowIndex, colIndex, &char, make([]string, 0)})
			} else {
				grid = append(grid, GridPosition{rowIndex, colIndex, nil, make([]string, 0)})
			}
		}
	}
	return antennaMap, grid, maxGridSize
}

func addAntinodes(antennaMap AntennaMap, grid []GridPosition, part2 bool) []GridPosition {
	for frequency, antennas := range antennaMap {
		for a := 0; a < len(antennas); a++ {
			antennaA := antennas[a]
			for b := a + 1; b < len(antennas); b++ {
				antennaB := antennas[b]

				distance := getPositionDistance(antennaA, antennaB)
				if part2 {
					antinodeAIndex := getGridIndex(antennaA, grid)
					if antinodeAIndex != -1 {
						grid[antinodeAIndex].antinodes = append(grid[antinodeAIndex].antinodes, frequency)
					}
					antinodeBIndex := getGridIndex(antennaB, grid)
					if antinodeBIndex != -1 {
						grid[antinodeBIndex].antinodes = append(grid[antinodeBIndex].antinodes, frequency)
					}
				}

				addAntinodesLoop(frequency, antennaA, distance, grid, "-", part2)

				addAntinodesLoop(frequency, antennaB, distance, grid, "+", part2)
			}
		}
	}
	return grid
}

func countAntinodes(grid []GridPosition) int {
	count := 0
	for _, cell := range grid {
		if len(cell.antinodes) > 0 {
			count++
		}
	}
	return count
}

func getPositionDistance(a MapPosition, b MapPosition) MapPosition {
	return MapPosition{b.row - a.row, b.col - a.col}
}

func getGridIndex(m MapPosition, grid []GridPosition) (index int) {
	index = slices.IndexFunc(grid, func(pos GridPosition) bool {
		if pos.row == m.row && pos.col == m.col {
			return true
		}
		return false
	})
	return index
}

func addAntinodesLoop(frequency string, startPos MapPosition, distance MapPosition, grid []GridPosition, operand string, part2 bool) {
	var antinode MapPosition
	if operand == "+" {
		antinode = MapPosition{startPos.row + distance.row, startPos.col + distance.col}
	} else if operand == "-" {
		antinode = MapPosition{startPos.row - distance.row, startPos.col - distance.col}
	}
	antinodeIndex := getGridIndex(antinode, grid)
	for antinodeIndex != -1 {
		grid[antinodeIndex].antinodes = append(grid[antinodeIndex].antinodes, frequency)
		if !part2 {
			return
		} else {
			if operand == "+" {
				antinode = MapPosition{antinode.row + distance.row, antinode.col + distance.col}
			} else if operand == "-" {
				antinode = MapPosition{antinode.row - distance.row, antinode.col - distance.col}
			}
			antinodeIndex = getGridIndex(antinode, grid)
		}
	}
}
