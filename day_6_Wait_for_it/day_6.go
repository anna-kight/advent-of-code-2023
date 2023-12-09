package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, fileScanner := getFileScanner("input.txt")
	var time, distance []int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, ":")
		digits := regexp.MustCompile("[0-9]+")
		if values[0] == "Time" {
			numbers := digits.FindAllString(values[1], -1)
			for i := 0; i < len(numbers); i++ {
				num, _ := strconv.Atoi(numbers[i])
				time = append(time, num)
			}
		} else if values[0] == "Distance" {
			numbers := digits.FindAllString(values[1], -1)
			for i := 0; i < len(numbers); i++ {
				num, _ := strconv.Atoi(numbers[i])
				distance = append(distance, num)
			}
		}
	}
	file.Close()

	var waysToWin []int
	for i := 0; i < len(time); i++ {
		winCount := 0
		for holdT := 0; holdT < time[i]; holdT++ {
			dist := (time[i] - holdT) * holdT
			if dist > distance[i] {
				winCount++
			}
		}
		waysToWin = append(waysToWin, winCount)
	}

	product := waysToWin[0] * waysToWin[1]
	for i := 2; i < len(waysToWin); i++ {
		product = product * waysToWin[i]
	}
	fmt.Println("Ways to win: ", fmt.Sprint(waysToWin))
	fmt.Println("The product of the number of ways to win in each race is: ", fmt.Sprint(product))
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
