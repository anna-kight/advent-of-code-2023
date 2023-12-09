package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type nextNode struct {
	L string
	R string
}

func main() {
	file, fileScanner := getFileScanner("input.txt")
	nodes := make(map[string]nextNode)
	var instructions string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !strings.Contains(line, "=") {
			instructions = line
		} else {
			split := strings.Split(line, " = ")
			key := split[0]
			next := strings.Trim(split[1], "()")
			nextSplit := strings.Split(next, ", ")
			val := nextNode{L: nextSplit[0], R: nextSplit[1]}
			nodes[key] = val
		}

	}
	file.Close()
	steps := 0
	currentNode := "AAA"
	i := 0
	for currentNode != "ZZZ" {
		if i >= len(instructions) {
			i = 0
		}
		if instructions[i] == 'L' {
			currentNode = nodes[currentNode].L
		} else {
			currentNode = nodes[currentNode].R
		}
		steps++
		i++
	}

	fmt.Println("The total number of steps is: ", fmt.Sprint(steps))
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
