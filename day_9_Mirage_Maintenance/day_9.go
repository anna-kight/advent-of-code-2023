package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, fileScanner := getFileScanner("input.txt")
	var extrapolated []int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")
		var intValues []int
		for _, val := range values {
			valInt, _ := strconv.Atoi(val)
			intValues = append(intValues, valInt)
		}
		extrapolated = append(extrapolated, findExtrap(intValues))
	}
	file.Close()

	sum := 0
	for _, val := range extrapolated {
		sum += val
	}

	fmt.Println("The sum of the extrapolated values is: ", fmt.Sprint(sum))
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

func findExtrap(known []int) (extrap int) {
	var sequences [][]int
	sequences = append(sequences, known)
	currentSeq := sequences[0]
	for !allZero(currentSeq) {
		var nextSeq []int
		for i := 0; i < len(currentSeq)-1; i++ {
			nextSeq = append(nextSeq, currentSeq[i+1]-currentSeq[i])
		}
		currentSeq = nextSeq
		sequences = append(sequences, currentSeq)
	}
	for i := len(sequences) - 1; i >= 1; i-- {
		sequences[i-1] = append(sequences[i-1], sequences[i][len(sequences[i])-1]+sequences[i-1][len(sequences[i-1])-1])
	}
	extrap = sequences[0][len(sequences[0])-1]

	return extrap
}

func allZero(check []int) bool {
	for _, val := range check {
		if val != 0 {
			return false
		}
	}
	return true
}
