package day01

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Day01() {
	// read a text file line by line

	file := "day01/01.txt"
	linesOrig := readLines(file)
	lines := parseSpelledDigits(linesOrig)

	lineCount := len(lines)

	// Find the first and the last digit of each input line

	first := make([]int, lineCount)
	last := make([]int, lineCount)

	for i, line := range lines {
		firstFound := false
		for _, char := range line {
			if unicode.IsDigit(char) {
				if !firstFound {
					first[i] = int(char - '0')
					firstFound = true
				}
				last[i] = int(char - '0')
			}
		}
	}

	results := []int{}

	for i := 0; i < len(first); i++ {
		str1 := strconv.Itoa(first[i])
		str2 := strconv.Itoa(last[i])

		resultStr := str1 + str2

		result, err := strconv.Atoi(resultStr)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
	}

	// Find the sum of the results

	sum := 0

	for _, result := range results {
		sum += result
	}

	// Print the sum

	println(sum)
}

// func Day01WithWords() {
// 	// read a text file line by line

// 	file := "day01/01.txt"
// 	lines := readLines(file)

// 	lineCount := len(lines)

// 	// Find the first and the last digit of each input line

// 	first := make([]int, lineCount)
// 	last := make([]int, lineCount)

// 	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// 	// replace the words with numbers

// 	for i, line := range lines {
// 	}
// }

func parseSpelledDigits(input []string) []string {
	// the words "one" to "nine" are spelled out in the input
	// parse them into digits 1 to 9 and replace them in the input

	replacer := strings.NewReplacer("one", "onee", "two", "twoo", "three", "threee", "four", "fourr", "five", "fivee", "six", "sixx", "seven", "sevenn", "eight", "eightt", "nine", "ninee")

	for i, line := range input {
		input[i] = replacer.Replace(line)
	}

	replacer = strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")

	for i, line := range input {
		input[i] = replacer.Replace(line)
	}

	return input
}

// readLines reads a whole file into memory
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
