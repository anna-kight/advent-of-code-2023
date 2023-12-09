package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards          string
	bid            int
	handType       int
	winnings       int
	numericalCards []int
}

type ByRank []hand

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	if a[i].handType != a[j].handType {
		return a[i].handType < a[j].handType

	} else {
		for k := 0; k <= 4; k++ {
			if a[i].numericalCards[k] != a[j].numericalCards[k] {
				return a[i].numericalCards[k] < a[j].numericalCards[k]
			}
		}
	}
	return a[i].handType < a[j].handType
}

func main() {
	file, fileScanner := getFileScanner("input.txt")
	var hands []hand
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")
		newHand := hand{cards: values[0]}
		newHand.bid, _ = strconv.Atoi(values[1])
		newHand.handType = getType(newHand.cards)
		newHand.numericalCards = getCardNums(newHand.cards)
		hands = append(hands, newHand)
	}
	file.Close()

	sort.Sort(ByRank(hands))

	totalWinnings := 0
	for i, currentHand := range hands {
		currentHand.winnings = (i + 1) * currentHand.bid
		totalWinnings += currentHand.winnings
	}

	fmt.Println("The total winnings are: ", fmt.Sprint(totalWinnings))
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

func getType(cards string) int {
	// hand types:
	// 7 Five of a kind
	// 6 Four of a kind
	// 5 Full house
	// 4 Three of a kind
	// 3 Two pair
	// 2 One pair
	// 1 High card
	counts := make(map[rune]int)
	for _, char := range cards {
		counts[char]++
	}
	numJs := 0
	val, ok := counts['J']
	if ok {
		numJs = val
		delete(counts, 'J')
	}
	if numJs == 5 {
		return 7
	}
	var justCount []int
	for _, count := range counts {
		justCount = append(justCount, count)
	}
	max := 0
	imax := -1
	for i, count := range justCount {
		if count > max {
			max = count
			imax = i
		}
	}
	justCount[imax] = justCount[imax] + numJs

	if slices.Contains(justCount, 5) {
		return 7
	} else if slices.Contains(justCount, 4) {
		return 6
	} else if slices.Contains(justCount, 3) {
		if slices.Contains(justCount, 2) {
			return 5
		} else {
			return 4
		}
	} else if slices.Contains(justCount, 2) {
		if len(justCount) == 3 {
			return 3
		} else {
			return 2
		}
	} else {
		return 1
	}
}

func getCardNums(cards string) (nums []int) {
	digits := regexp.MustCompile("[0-9]+")
	var charNum int
	for _, slice := range strings.Split(cards, "") {
		if digits.MatchString(slice) {
			charNum, _ = strconv.Atoi(slice)
		} else if slice == "T" {
			charNum = 10
		} else if slice == "J" {
			charNum = 1
		} else if slice == "Q" {
			charNum = 11
		} else if slice == "K" {
			charNum = 12
		} else if slice == "A" {
			charNum = 13
		}
		nums = append(nums, charNum)
	}

	return nums
}
