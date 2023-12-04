// Determine which games would have been possible if the bag had been loaded with only 12 red cubes,
// 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id       int
	possible bool
	power    int
}

func main() {
	file, fileScanner := getFileScanner("input.txt")

	sumOfPossibleGames := 0
	sumOfPowers := 0
	var currentGame game
	for fileScanner.Scan() {
		line := fileScanner.Text()
		currentGame = lineToGame(line)

		if currentGame.possible {
			sumOfPossibleGames += currentGame.id
		}
		sumOfPowers += currentGame.power
	}
	fmt.Println("The sum of all of the possible games is: ", fmt.Sprint(sumOfPossibleGames))
	fmt.Println("The sum of the powers is: ", fmt.Sprint(sumOfPowers))

	file.Close()
}

func getFileScanner(fileName string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	return file, fileScanner
}

func lineToGame(line string) (newGame game) {
	splitLine := strings.Split(line, ":")
	newGame.id, _ = strconv.Atoi(strings.Trim(splitLine[0], "Game "))
	rounds := strings.Split(splitLine[1], ";")
	newGame.power = getPower(rounds)
	for i := 0; i < len(rounds); i++ {
		colors := strings.Split(rounds[i], ",")
		for j := 0; j < len(colors); j++ {
			if !validColorAmount(colors[j]) {
				newGame.possible = false
				return newGame
			}
		}
	}
	newGame.possible = true
	return newGame
}

func validColorAmount(colorAmount string) bool {
	split := strings.Split(colorAmount, " ")
	number, _ := strconv.Atoi(split[1])
	switch color := split[2]; color {
	case "red":
		if number > 12 {
			return false
		}
	case "green":
		if number > 13 {
			return false
		}
	case "blue":
		if number > 14 {
			return false
		}
	default:
		return false
	}
	return true
}

func getPower(colors []string) int {
	minRed, minGreen, minBlue := 0, 0, 0
	for i := 0; i < len(colors); i++ {
		colorAmounts := strings.Split(colors[i], ",")
		for j := 0; j < len(colorAmounts); j++ {
			split := strings.Split(colorAmounts[j], " ")
			number, _ := strconv.Atoi(split[1])
			switch color := split[2]; color {
			case "red":
				if minRed < number {
					minRed = number
				}
			case "green":
				if minGreen < number {
					minGreen = number
				}
			case "blue":
				if minBlue < number {
					minBlue = number
				}
			default:
				return 0
			}
		}
	}
	return minRed * minBlue * minGreen
}
