package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	number     int
	count      int
	points     int
	matchCount int
}

func main() {
	file, fileScanner := getFileScanner("input.txt")

	totalPoints := 0
	var cards []card
	for fileScanner.Scan() {
		cardData := fileScanner.Text()
		number, points, matches := calculatePoints(cardData)
		newCard := card{
			number:     number,
			count:      1,
			points:     points,
			matchCount: matches,
		}
		cards = append(cards, newCard)

		totalPoints += points
	}
	fmt.Println("The point total for all the cards is: ", fmt.Sprint(totalPoints))

	totalCards := 0
	for i, current := range cards {
		for n := 1; n <= current.matchCount; n++ {
			cards[i+n].count += current.count
		}
	}
	for _, card := range cards {
		totalCards += card.count
	}

	fmt.Println("The total number of scratchcards is: ", fmt.Sprint(totalCards))

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

func calculatePoints(card string) (cNumber int, points int, matches int) {
	points = 0
	matches = 0
	digits := regexp.MustCompile("[0-9]+")
	splitCardNum := strings.Split(card, ":")
	cNumber, _ = strconv.Atoi(strings.Trim(splitCardNum[0], "Card "))
	numbers := strings.Split(splitCardNum[1], "|")
	winningNumbers := digits.FindAllString(numbers[0], -1)
	numbersElfHas := digits.FindAllString(numbers[1], -1)
	for _, elfNum := range numbersElfHas {
		if isWinner(winningNumbers, elfNum) {
			matches++
			if points == 0 {
				points = 1
			} else {
				points = 2 * points
			}
		}
	}
	return cNumber, points, matches
}

func isWinner(winNumbers []string, have string) bool {
	for _, val := range winNumbers {
		if val == have {
			return true
		}
	}
	return false
}
