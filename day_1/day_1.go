package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, fileScanner := getFileScanner("input.txt")

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		calibrationValue, err := getCalibrationValue(line)
		if err != nil {
			fmt.Println(err)
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
	var firstDigit, secondDigit string
	for i := 0; i < len(line); i++ {
		if isNumber(string(line[i])) {
			firstDigit = string(line[i])
			fmt.Println("The first digit is: ", fmt.Sprint(firstDigit))
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if isNumber(string(line[i])) {
			secondDigit = string(line[i])
			fmt.Println("The second digit is: ", fmt.Sprint(secondDigit))
			break
		}
	}
	return strconv.Atoi(firstDigit + secondDigit)
}
