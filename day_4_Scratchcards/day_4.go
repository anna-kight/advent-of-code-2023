package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, fileScanner := getFileScanner("input.txt")

	totalPoints := 0
	for fileScanner.Scan() {
		card := fileScanner.Text()
		points := calculatePoints(card)

		totalPoints += points
	}
	fmt.Println("The point total for all the cards is: ", fmt.Sprint(totalPoints))

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

func calculatePoints(card string) (points int) {
	points = 0
	digits := regexp.MustCompile("[0-9]+")
	splitCardNum := strings.Split(card, ":")
	numbers := strings.Split(splitCardNum[1], "|")
	winningNumbers := digits.FindAllString(numbers[0], -1)
	numbersElfHas := digits.FindAllString(numbers[1], -1)
	for _, elfNum := range numbersElfHas {
		if isWinner(winningNumbers, elfNum) {
			if points == 0 {
				points = 1
			} else {
				points = 2 * points
			}
		}
	}
	return points
}

func isWinner(winNumbers []string, have string) bool {
	for _, val := range winNumbers {
		if val == have {
			return true
		}
	}
	return false
}
