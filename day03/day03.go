package day03

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func Day03() {
	input := readLines("day03/03.txt")

	grid := make([][]rune, 140)

	sum := 0

	// initialize grid
	for i := range grid {
		grid[i] = make([]rune, 140)
	}

	for i, line := range input {
		// loop over each character in line
		for j, char := range line {
			grid[i][j] = char
		}
	}

	// Search for numbers. Numbers can have multiple digits. If the number adjecent to a symbol (not a .), add it to the sum.

	for i, line := range grid {
		for j := 0; j < len(line); j++ {
			char := grid[i][j]
			if unicode.IsDigit(char) {
				if findPunctArround(i, j, grid) {
					sum = sum + getCurrentNumber(i, j, grid)
					if unicode.IsDigit(grid[i][j+1]) {
						j++
						if unicode.IsDigit(grid[i][j+1]) {
							j++
						}
					}
				}
			}
		}
	}
	println(sum)
}

func findPunctArround(i int, j int, grid [][]rune) bool {
	// check for i+1, i-1, j+1, j-1, i+1j+1, i+1j-1, i-1j+1, i-1j-1
	if i+1 < len(grid) {
		if isPunct(grid[i+1][j]) {
			return true
		}
	}
	if i-1 >= 0 {
		if isPunct(grid[i-1][j]) {
			return true
		}
	}
	if j+1 < len(grid[i]) {
		if isPunct(grid[i][j+1]) {
			return true
		}
	}
	if j-1 >= 0 {
		if isPunct(grid[i][j-1]) {
			return true
		}
	}
	if i+1 < len(grid) && j+1 < len(grid[i]) {
		if isPunct(grid[i+1][j+1]) {
			return true
		}
	}
	if i+1 < len(grid) && j-1 >= 0 {
		if isPunct(grid[i+1][j-1]) {
			return true
		}
	}
	if i-1 >= 0 && j+1 < len(grid[i]) {
		if isPunct(grid[i-1][j+1]) {
			return true
		}
	}
	if i-1 >= 0 && j-1 >= 0 {
		if isPunct(grid[i-1][j-1]) {
			return true
		}
	}
	return false
}

func isPunct(char rune) bool {
	if char == '@' || char == '&' || char == '$' || char == '%' || char == '=' || char == '+' || char == '-' || char == '*' || char == '#' || char == '/' {
		if char != '.' {
			return true
		}
	}
	return false
}

func getCurrentNumber(i int, j int, grid [][]rune) int {
	numberString := string(grid[i][j])

	if unicode.IsDigit(grid[i][j+1]) {
		numberString = numberString + string(grid[i][j+1])
		if unicode.IsDigit(grid[i][j+2]) {
			numberString = numberString + string(grid[i][j+2])
		} else if unicode.IsDigit(grid[i][j-1]) {
			numberString = string(grid[i][j-1]) + numberString
		}
	} else if unicode.IsDigit(grid[i][j-1]) {
		numberString = string(grid[i][j-1]) + numberString
		if unicode.IsDigit(grid[i][j-2]) {
			numberString = string(grid[i][j-2]) + numberString
		} else if unicode.IsDigit(grid[i][j+1]) {
			numberString = numberString + string(grid[i][j+1])
		}
	}

	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}
	println(number)
	return number
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); scanner.Err() != nil {
		panic(err)
	}
	return lines
}
