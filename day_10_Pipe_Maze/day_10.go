package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type pipe struct {
	pipeType    string
	location    [2]int
	connections [][2]int
}

const Pi = 3.14

var connectToUp = [3]string{"|", "7", "F"}
var connectToDown = [3]string{"|", "L", "7"}
var connectToRight = [3]string{"-", "J", "7"}
var connectToLeft = [3]string{"-", "L", "F"}

func main() {
	file, fileScanner := getFileScanner("input.txt")
	var pipeGrid [][]pipe
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var pipeLine []pipe
		for c := 0; c < len(line); c++ {
			pipeLine = append(pipeLine, pipe{pipeType: string(line[c]), location: [2]int{i, c}})
			pipeLine[c].connections = findConnections(pipeLine[c])
		}
		pipeGrid = append(pipeGrid, pipeLine)
		i++
	}
	file.Close()
	var seenLocations [][2]int
	currentPipe := findS(pipeGrid)

	for _, c := range currentPipe.connections {
		if slices.Contains(pipeGrid[c[0]][c[1]].connections, currentPipe.location) {
			seenLocations = append(seenLocations, currentPipe.location)
			currentPipe = pipeGrid[c[0]][c[1]]
			break
		}
	}

	moved := true
	for moved {
		moved = false
		for _, c := range currentPipe.connections {
			if !slices.Contains(seenLocations, pipeGrid[c[0]][c[1]].location) {
				seenLocations = append(seenLocations, currentPipe.location)
				currentPipe = pipeGrid[c[0]][c[1]]
				moved = true
				// fmt.Println(currentPipe)
				break
			}
		}
	}

	lengthToFarthest := (len(seenLocations) + 1) / 2

	fmt.Println("The number of steps to get from the starting position to the farthest point is: ", fmt.Sprint(lengthToFarthest))
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

func findS(grid [][]pipe) (location pipe) {
	for _, r := range grid {
		for _, p := range r {
			if p.pipeType == "S" {
				location = p
			}
		}
	}
	return location
}

func findConnections(p pipe) (connections [][2]int) {
	up, down, left, right := getAdjacent(p.location)
	var potentialConnections [][2]int
	switch p.pipeType {
	case "S":
		potentialConnections = [][2]int{up, down, left, right}
	case "|":
		potentialConnections = [][2]int{up, down}
	case "-":
		potentialConnections = [][2]int{left, right}
	case "L":
		potentialConnections = [][2]int{up, right}
	case "J":
		potentialConnections = [][2]int{up, left}
	case "7":
		potentialConnections = [][2]int{down, left}
	case "F":
		potentialConnections = [][2]int{down, right}
	default:
		potentialConnections = [][2]int{}
	}
	for _, c := range potentialConnections {
		if c[0] == -1 || c[1] == -1 {
			continue
		}
		connections = append(connections, c)
	}

	return connections
}

func getAdjacent(p [2]int) (up [2]int, down [2]int, left [2]int, right [2]int) {
	up, down, left, right = [2]int{-1, -1}, [2]int{-1, -1}, [2]int{-1, -1}, [2]int{-1, -1}
	if p[0] > 0 {
		up = [2]int{p[0] - 1, p[1]}
	}
	down = [2]int{p[0] + 1, p[1]}
	if p[1] > 0 {
		left = [2]int{p[0], p[1] - 1}
	}
	right = [2]int{p[0], p[1] + 1}

	return up, down, left, right
}
