package day04

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day04() {
	input := readLines("day04/04.txt")

	part01(input)
	part02(input)
}

func part01(input []string) {
	sum := 0

	for _, line := range input {
		// remove first 8 chars
		line = line[8:]
		// split line at |
		lists := strings.Split(line, "|")

		var numbersWinning []int
		var numbersHave []int

		// loop over each list
		for i, list := range lists {
			// use strings.Fields to split at spaces
			numbers := strings.Fields(list)
			// loop over each number
			for _, number := range numbers {
				// convert to int
				numberInt, _ := strconv.Atoi(number)
				// check if number is in winning list
				if i == 0 {
					numbersWinning = append(numbersWinning, numberInt)
				}
				// check if number is in have list
				if i == 1 {
					numbersHave = append(numbersHave, numberInt)
				}
			}
		}

		// check how many numbers are in both lists
		var count int
		for _, number := range numbersWinning {
			for _, number2 := range numbersHave {
				if number == number2 {
					count++
				}
			}
		}
		// Berechne 2^(count-1)
		sum += int(math.Pow(2, float64(count-1)))
	}
	println(sum)
}

func part02(input []string) {
	cardCount := make([]int, len(input))

	for j, line := range input {
		cardCount[j]++
		// remove first 8 chars
		line = line[8:]
		// split line at |
		lists := strings.Split(line, "|")

		var numbersWinning []int
		var numbersHave []int

		// loop over each list
		for i, list := range lists {
			// use strings.Fields to split at spaces
			numbers := strings.Fields(list)
			// loop over each number
			for _, number := range numbers {
				// convert to int
				numberInt, _ := strconv.Atoi(number)
				// check if number is in winning list
				if i == 0 {
					numbersWinning = append(numbersWinning, numberInt)
				}
				// check if number is in have list
				if i == 1 {
					numbersHave = append(numbersHave, numberInt)
				}
			}
		}

		// check how many numbers are in both lists
		var matchingNumbers int
		for _, number := range numbersWinning {
			for _, number2 := range numbersHave {
				if number == number2 {
					matchingNumbers++
				}
			}
		}

		for count := 0; count < matchingNumbers; count++ {
			for x := 0; x < cardCount[j]; x++ {
				cardCount[j+count+1]++
			}
		}
	}

	// sum up all card counts
	var sum int
	for _, count := range cardCount {
		sum += count
	}
	println(sum)
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
