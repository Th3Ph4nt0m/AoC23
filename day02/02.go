package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day02part01() {
	lines := readLines("day02/02.txt")

	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

	sum1 := 0

	for _, line := range lines {
		// split line at :
		firstSplit := strings.Split(line, ":")
		gameIdString := firstSplit[0]
		gameId, err := strconv.Atoi(strings.Split(gameIdString, " ")[1])
		if err != nil {
			panic(err)
		}

		println("Start of game ", gameId)

		possible := true

		// split line at ;
		roundsString := firstSplit[1]

		roundsStringSplit := strings.Split(roundsString, ";")
		for _, roundString := range roundsStringSplit {
			// split line at ,
			colorStringSplit := strings.Split(roundString, ",")

			for _, colorString := range colorStringSplit {
				colorStringSplit2 := strings.Split(colorString, " ")
				colorCount, err := strconv.Atoi(colorStringSplit2[1])
				if err != nil {
					panic(err)
				}
				color := colorStringSplit2[2]

				if color == "red" {
					println("Count of red: ", colorCount)
					if colorCount > 12 {
						possible = false
					}
					println("Possible: ", possible)
				}
				if color == "blue" {
					println("Count of blue: ", colorCount)
					if colorCount > 14 {
						possible = false
					}
					println("Possible: ", possible)
				}
				if color == "green" {
					println("Count of green: ", colorCount)
					if colorCount > 13 {
						possible = false
					}
					println("Possible: ", possible)
				}
			}
		}
		if possible {
			sum1 += gameId
			println("!!!!!!Added ", gameId, " to sum")
		}
	}

	println(sum1)
}

func Day02part02() {
	lines := readLines("day02/02.txt")

	sum := 0

	for _, line := range lines {
		game := strings.TrimSpace(strings.Split(line, ":")[1])
		colorsOnly := strings.Replace(game, ";", ",", -1)

		colors := strings.Split(colorsOnly, ",")
		redMin := 1
		blueMin := 1
		greenMin := 1

		for _, color := range colors {
			colorSplit := strings.Split(strings.TrimSpace(color), " ")
			colorCount, err := strconv.Atoi(colorSplit[0])
			if err != nil {
				println("!!!!!!!!", color)
				panic(err)
			}
			colorName := colorSplit[1]

			if colorName == "red" {
				if colorCount > redMin {
					redMin = colorCount
				}
			}
			if colorName == "blue" {
				if colorCount > blueMin {
					blueMin = colorCount
				}
			}
			if colorName == "green" {
				if colorCount > greenMin {
					greenMin = colorCount
				}
			}
		}
		power := redMin * blueMin * greenMin
		sum += power
	}
	println(sum)
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
