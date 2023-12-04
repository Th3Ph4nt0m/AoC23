package day03

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func Day03() {
	input := readLines("day03/day_03.txt")

	grid := make([][]rune, 140)

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

	println("Part 01: ", part01(grid))
	println("Part 02: ", part02(grid))
}

func part01(grid [][]rune) int {
	sum := 0
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
	return sum
}

func part02(grid [][]rune) int {
	sum := 0
	for i, line := range grid {
		for j := 0; j < len(line); j++ {
			char := grid[i][j]
			if char == '*' {
				ratio := findNumbersArround(i, j, grid)
				if ratio != -1 {
					sum += ratio
				}
			}
		}
	}
	return sum
}

func checkIfFull(slice []int) bool {
	if (slice[0] == -1) && (slice[1] == -1) {
		return false
	} else if (slice[0] == -1) || (slice[1] == -1) {
		return false
	}
	return true
}

func findNumbersArround(i int, j int, grid [][]rune) int {
	numbers := make([]int, 2)
	// init
	numbers[0] = -1
	numbers[1] = -1

	// check for i+1, i-1, j+1, j-1, i+1j+1, i+1j-1, i-1j+1, i-1j-1
	if i+1 < len(grid) {
		if unicode.IsDigit(grid[i+1][j]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i+1, j, grid)
				} else {
					if numbers[0] != getCurrentNumber(i+1, j, grid) {
						numbers[1] = getCurrentNumber(i+1, j, grid)
					}
				}
			}
		}
	}
	if i-1 >= 0 {
		if unicode.IsDigit(grid[i-1][j]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i-1, j, grid)
				} else {
					if numbers[0] != getCurrentNumber(i-1, j, grid) {
						numbers[1] = getCurrentNumber(i-1, j, grid)
					}
				}
			}
		}
	}
	if j+1 < len(grid[i]) {
		if unicode.IsDigit(grid[i][j+1]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i, j+1, grid)
				} else {
					if numbers[0] != getCurrentNumber(i, j+1, grid) {
						numbers[1] = getCurrentNumber(i, j+1, grid)
					}
				}
			}
		}
	}
	if j-1 >= 0 {
		if unicode.IsDigit(grid[i][j-1]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i, j-1, grid)
				} else {
					if numbers[0] != getCurrentNumber(i, j-1, grid) {
						numbers[1] = getCurrentNumber(i, j-1, grid)
					}
				}
			}
		}
	}
	if i+1 < len(grid) && j+1 < len(grid[i]) {
		if unicode.IsDigit(grid[i+1][j+1]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i+1, j+1, grid)
				} else {
					if numbers[0] != getCurrentNumber(i+1, j+1, grid) {
						numbers[1] = getCurrentNumber(i+1, j+1, grid)
					}
				}
			}
		}
	}
	if i+1 < len(grid) && j-1 >= 0 {
		if unicode.IsDigit(grid[i+1][j-1]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i+1, j-1, grid)
				} else {
					if numbers[0] != getCurrentNumber(i+1, j-1, grid) {
						numbers[1] = getCurrentNumber(i+1, j-1, grid)
					}
				}
			}
		}
	}
	if i-1 >= 0 && j+1 < len(grid[i]) {
		if unicode.IsDigit(grid[i-1][j+1]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i-1, j+1, grid)
				} else {
					if numbers[0] != getCurrentNumber(i-1, j+1, grid) {
						numbers[1] = getCurrentNumber(i-1, j+1, grid)
					}
				}
			}
		}
	}
	if i-1 >= 0 && j-1 >= 0 {
		if unicode.IsDigit(grid[i-1][j-1]) {
			if !checkIfFull(numbers) {
				if numbers[0] == -1 {
					numbers[0] = getCurrentNumber(i-1, j-1, grid)
				} else {
					if numbers[0] != getCurrentNumber(i-1, j-1, grid) {
						numbers[1] = getCurrentNumber(i-1, j-1, grid)
					}
				}
			}
		}
	}
	if numbers[0] != -1 && numbers[1] != -1 {
		return numbers[0] * numbers[1]
	}
	return -1
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
