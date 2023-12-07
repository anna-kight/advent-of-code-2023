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

func main() {

	var almanac []almanacMap
	var currentMap almanacMap
	var seeds []int
	digits := regexp.MustCompile("[0-9]+")
	file, fileScanner := getFileScanner("example.txt")
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Index(line, "seeds:") == 0 {
			seeds = getSeeds(line)
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
	file.Close()
	var lowestLocation int

	for _, seed := range seeds {
		location := getLocation(almanac, seed)
		if location < lowestLocation {
			lowestLocation = location
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

func getSeeds(line string) (nums []int) {
	seeds := strings.Split(line, " ")
	seeds = seeds[1:]
	for index, seed := range seeds {
		nums[index], _ = strconv.Atoi(seed)
	}
	return nums
}

func lineToMap(line string) (rangeMap [3]int) {
	numStrings := strings.Split(line, " ")
	for index, num := range numStrings {
		rangeMap[index], _ = strconv.Atoi(num)
	}
	return rangeMap
}

func getLocation(almanac []almanacMap, seed int) (location int) {
	return location
}
