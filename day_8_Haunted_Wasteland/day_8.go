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
		if !strings.Contains(line, "=") && strings.Contains(line, "L") {
			instructions = line
		} else if strings.Contains(line, "=") {
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
	currentNodes := getStartingNodes(nodes)
	i := 0
	for !allNodesEndInZ(currentNodes) {
		if i >= len(instructions) {
			i = 0
		}
		currentNodes = getNextNodes(currentNodes, nodes, instructions[i])
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

func getStartingNodes(nodeMap map[string]nextNode) (sNodes []string) {
	for name, _ := range nodeMap {
		if name[len(name)-1:] == "A" {
			sNodes = append(sNodes, name)
		}
	}
	return sNodes
}

func allNodesEndInZ(nodes []string) bool {
	for _, n := range nodes {
		if n[len(n)-1:] != "Z" {
			return false
		}
	}
	return true
}

func getNextNodes(nodes []string, nodeMap map[string]nextNode, direction byte) (nNodes []string) {
	nNodes = nodes
	for i, n := range nodes {
		if direction == 'L' {
			nNodes[i] = nodeMap[n].L
		} else {
			nNodes[i] = nodeMap[n].R
		}
	}
	return nNodes
}
