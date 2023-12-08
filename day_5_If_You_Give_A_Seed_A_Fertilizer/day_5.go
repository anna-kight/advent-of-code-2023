package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type almanacMap struct {
	sourceCat      string
	destinationCat string
	conversions    [][3]int
}

//69841804 too high

func main() {

	var almanac []almanacMap
	var currentMap almanacMap
	var seedRanges []int
	digits := regexp.MustCompile("[0-9]+")
	file, fileScanner := getFileScanner("input.txt")
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Index(line, "seeds:") == 0 {
			seedRanges = getSeedRanges(line)
		} else if strings.Contains(line, ":") {
			if currentMap.sourceCat != "" {
				almanac = append(almanac, currentMap)
			}
			line = strings.Trim(line, " map:")
			names := strings.Split(line, "-")
			currentMap = almanacMap{sourceCat: names[0], destinationCat: names[2]}
		} else if digits.MatchString(line) {
			newConversion := lineToMap(line)
			currentMap.conversions = append(currentMap.conversions, newConversion)
		}
	}
	almanac = append(almanac, currentMap)
	file.Close()
	lowestLocation := -2

	for i := 0; i < len(seedRanges)-1; i = i + 2 {
		seeds := getSeeds(seedRanges[i : i+2])

		for _, seed := range seeds {
			location := getLocation(almanac, seed)
			if location < lowestLocation || lowestLocation == -2 {
				lowestLocation = location
			}
		}
	}

	fmt.Println("The lowest location number that corresponds to one of the initial seed numbers is: ", fmt.Sprint(lowestLocation))
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

func getSeeds(seedRange []int) (nums []int) {
	// seeds := strings.Split(line, " ")
	// seeds = seeds[1:]
	// var seedRange []int
	// for _, seed := range seeds {
	// 	temp, _ := strconv.Atoi(seed)
	// 	seedRange = append(seedRange, temp)
	// }
	for i := 0; i < len(seedRange)-1; i = i + 2 {
		for j := 0; j < seedRange[i+1]; j++ {
			nums = append(nums, seedRange[i]+j)
		}
	}

	return nums
}

func getSeedRanges(line string) (seedRange []int) {
	seeds := strings.Split(line, " ")
	seeds = seeds[1:]
	for _, seed := range seeds {
		temp, _ := strconv.Atoi(seed)
		seedRange = append(seedRange, temp)
	}
	return seedRange
}

func lineToMap(line string) (rangeMap [3]int) {
	numStrings := strings.Split(line, " ")
	for index, num := range numStrings {
		rangeMap[index], _ = strconv.Atoi(num)
	}
	return rangeMap
}

func getLocation(almanac []almanacMap, seed int) (location int) {
	cat := "seed"
	num := seed
	for i := 0; i < 7; i++ {
		amap := findNextMap(almanac, cat)
		num = sourceToDest(num, amap)
		cat = amap.destinationCat
	}
	if cat == "location" {
		location = num
	} else {
		location = -1
	}

	return location
}

func findNextMap(almanac []almanacMap, cat string) almanacMap {
	for _, amap := range almanac {
		if cat == amap.sourceCat {
			return amap
		}
	}
	return almanacMap{}
}

func sourceToDest(num int, amap almanacMap) int {
	for _, con := range amap.conversions {
		if con[1] <= num && num <= con[1]+con[2] {
			offset := num - con[1]
			return con[0] + offset
		}
	}
	return num
}
