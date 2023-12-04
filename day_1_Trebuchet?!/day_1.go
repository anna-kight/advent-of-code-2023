package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numberStrings = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {

	file, fileScanner := getFileScanner("input.txt")

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		calibrationValue, err := getCalibrationValue(line)
		if err != nil {
			fmt.Println("Error getting calibration value: ", err)
		}
		sum += calibrationValue
	}
	fmt.Println("The sum of all of the calibration values is: ", fmt.Sprint(sum))

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

func isNumber(str string) bool {
	if _, err := strconv.Atoi(string(str)); err == nil {
		return true
	} else {
		return false
	}
}

func getCalibrationValue(line string) (int, error) {
	firstDigit, firstDigitIndex := getFirstDigit(line)
	lastDigit, lastDigitIndex := getLastDigit(line)

	firstNumber, firstNumberIndex := getFirstSpelledNumber(line)
	lastNumber, lastNumberIndex := getLastSpelledNumber(line)

	var first, last string
	if firstDigitIndex < firstNumberIndex || firstNumberIndex == -1 {
		first = firstDigit
	} else {
		first = firstNumber
	}

	if lastDigitIndex > lastNumberIndex || lastNumberIndex == -1 {
		last = lastDigit
	} else {
		last = lastNumber
	}
	return strconv.Atoi(first + last)
}

func getFirstDigit(line string) (digit string, index int) {
	for i := 0; i < len(line); i++ {
		if isNumber(string(line[i])) {
			return string(line[i]), i
		}
	}
	return
}

func getLastDigit(line string) (digit string, index int) {
	for i := len(line) - 1; i >= 0; i-- {
		if isNumber(string(line[i])) {
			return string(line[i]), i
		}
	}
	return
}

func getFirstSpelledNumber(line string) (number string, index int) {
	index = -1
	for i := 0; i < len(numberStrings); i++ {
		indexOfI := strings.Index(line, numberStrings[i])
		if indexOfI != -1 {
			if index > -1 && indexOfI < index {
				index = indexOfI
				number = strconv.Itoa(i + 1)
			} else if index == -1 {
				index = indexOfI
				number = strconv.Itoa(i + 1)
			}
		}
	}
	return number, index
}

func getLastSpelledNumber(line string) (number string, index int) {
	index = -1
	for i := len(numberStrings) - 1; i >= 0; i-- {
		indexOfI := strings.LastIndex(line, numberStrings[i])
		if indexOfI != -1 && indexOfI > index {
			index = indexOfI
			number = strconv.Itoa(i + 1)
		}
	}
	return number, index
}
