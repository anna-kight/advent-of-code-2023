package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

type Galaxy struct {
	row int
	col int
}

func main() {
	file, fileScanner := getFileScanner("input.txt")

	var unexpanded []Galaxy
	row := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for currentCol, c := range line {
			if string(c) == "#" {
				unexpanded = append(unexpanded, Galaxy{col: currentCol, row: row})
			}
		}
		row++
	}
	expanded := expand(unexpanded)

	var pairs [][2]Galaxy
	for i, g := range expanded {
		for j := i + 1; j < len(expanded); j++ {
			pairs = append(pairs, [2]Galaxy{g, expanded[j]})
		}
	}
	sumOfPaths := 0
	for _, p := range pairs {
		path := findShortestPath(p)
		sumOfPaths += path
	}

	fmt.Println("The sum of the shotrest paths is: ", fmt.Sprint(sumOfPaths))

	fmt.Print()
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

func expand(u []Galaxy) (expanded []Galaxy) {
	var rows, cols []int
	for _, g := range u {
		expanded = append(expanded, g)
		rows = append(rows, g.row)
		cols = append(cols, g.col)
	}
	maxRow := slices.Max(rows)
	slices.Sort(rows)
	slices.Reverse(rows)
	maxCol := slices.Max(cols)
	slices.Sort(cols)
	slices.Reverse(cols)
	for r := maxRow; r >= 0; r-- {
		if !slices.Contains(rows, r) {
			for i, g := range expanded {
				if g.row > r {
					expanded[i].row = expanded[i].row + 999999
				}
			}
		}
	}
	for c := maxCol; c >= 0; c-- {
		if !slices.Contains(cols, c) {
			for i, g := range expanded {
				if g.col > c {
					expanded[i].col = expanded[i].col + 999999
				}
			}
		}
	}
	return expanded
}

func findShortestPath(pair [2]Galaxy) (shortestPath int) {
	shortestPath = int(math.Abs(float64(pair[0].col)-float64(pair[1].col)) + math.Abs(math.Abs(float64(pair[0].row)-float64(pair[1].row))))
	return shortestPath
}
