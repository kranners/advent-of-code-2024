package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const XMAS = "XMAS"

type Vector struct {
	Row int
	Col int
}

func move(position Vector, direction Vector) Vector {
	return Vector{Row: position.Row + direction.Row, Col: position.Col + direction.Col}
}

func at(lines []string, position Vector) (letter byte, outOfBounds error) {
	if position.Row < 0 || position.Row >= len(lines) {
		var nothing byte
		return nothing, fmt.Errorf("row out of bounds")
	}

	if position.Col < 0 || position.Col >= len(lines[position.Row]) {
		var nothing byte
		return nothing, fmt.Errorf("col out of bounds")
	}

	return lines[position.Row][position.Col], nil
}

func readWordInDirection(lines []string, position Vector, direction Vector, indexInWord int) bool {
	letter, err := at(lines, position)

	if err != nil {
		return false
	}

	if letter != XMAS[indexInWord] {
		return false
	}

	if indexInWord == len(XMAS)-1 {
		return true
	}

	position = move(position, direction)
	return readWordInDirection(lines, position, direction, indexInWord+1)
}

func readWordAtPosition(lines []string, position Vector) int {
	directions := [8]Vector{
		{-1, -1},
		{0, -1},
		{1, -1},

		{-1, 0},
		{1, 0},

		{-1, 1},
		{0, 1},
		{1, 1},
	}

	matches := 0
	for _, direction := range directions {
		if readWordInDirection(lines, position, direction, 0) {
			matches += 1
		}
	}

	return matches
}

func checkCornersAroundPosition(lines []string, position Vector) bool {
	var corners []byte
	directions := [4]Vector{
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}

	for _, direction := range directions {
		corner, err := at(lines, move(position, direction))

		if err != nil {
			return false
		}

		if corner != 'M' && corner != 'S' {
			return false
		}

		corners = append(corners, corner)
	}

	if corners[0] == corners[1] || corners[2] == corners[3] {
		return false
	}

	return true
}

func readXMasAtPosition(lines []string, position Vector) bool {
	letter, err := at(lines, position)

	if err != nil {
		return false
	}

	if letter != 'A' {
		return false
	}

	return checkCornersAroundPosition(lines, position)
}

func main() {
	input, err := os.ReadFile("./input.txt")
	check(err)

	matches := 0

	lines := strings.Split(string(input), "\n")
	for row, line := range lines {
		for col := range line {
			position := Vector{row, col}

			if readXMasAtPosition(lines, position) {
				matches += 1
			}
		}
	}

	fmt.Println(matches)
}
