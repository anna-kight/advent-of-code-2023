package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 528288 too low, 549924 too high
func main() {
	file, fileScanner := getFileScanner("input.txt")

	sumOfPartNumbers := 0
	var data []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		data = append(data, line)
	}
	file.Close()

	for i := 0; i < len(data); i++ {
		digits := regexp.MustCompile("[0-9]+")
		numbers := digits.FindAllString(data[i], -1)
		for j := 0; j < len(numbers); j++ {
			index := strings.Index(data[i], numbers[j])
			symbols := regexp.MustCompile("[^0-9|.]")
			partNumber, _ := strconv.Atoi(numbers[j])
			numLen := len(numbers[j])
			lower := index - 1
			upper := index + numLen + 1
			if lower < 0 {
				lower = 0
			}
			if upper >= len(data) {
				upper = len(data) - 1
			}

			isPart := false
			// under := data[i+1][lower:upper]
			// over := data[i-1][lower:upper]
			if (symbols.MatchString(string(data[i][lower]))) || (index < len(data[i])-1 && symbols.MatchString(string(data[i][upper-1]))) {
				isPart = true
			} else if i > 0 && symbols.MatchString(string(data[i-1][lower:upper])) {
				isPart = true
			} else if i < len(data)-1 && symbols.MatchString(string(data[i+1][lower:upper])) {
				isPart = true
			}
			if isPart {
				sumOfPartNumbers += partNumber
			}

		}

	}

	fmt.Println("The sum of all of the part numbers is: ", sumOfPartNumbers)

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

func isNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	} else {
		return false
	}
}
