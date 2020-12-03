package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	tree      byte = '#'
	openSpace byte = '.'
	stepDown  int  = 2
	stepRight int  = 1
)

func main() {
	dayThree()
}
func dayThree() {
	data, err := ioutil.ReadFile("inputs/day3.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	temp := strings.Split(string(data), "\n")

	treeMap := make([][]byte, len(temp))

	for lineNum := 0; lineNum < len(temp); lineNum++ {
		line := temp[lineNum]

		newSlice := make([]byte, len(line))
		treeMap[lineNum] = newSlice

		for i := 0; i < len(line); i++ {
			treeMap[lineNum][i] = line[i]
		}
	}
	currentStepDown := 0
	currentStepRight := 0
	lenghtOfTrack := len(treeMap[0])
	numberOfTreesHit := 0
	for currentStepDown < (len(treeMap) - 1) {
		currentStepDown += stepDown
		currentStepRight += stepRight
		if currentStepRight >= lenghtOfTrack {
			currentStepRight -= lenghtOfTrack
		}
		if currentStepDown >= len(treeMap) {
			break
		}
		currentPos := treeMap[currentStepDown][currentStepRight]
		if currentPos == tree {
			numberOfTreesHit++
		}
	}

	fmt.Printf("It hit %d trees\n", numberOfTreesHit)

}

func dayTwo() {
	data, err := ioutil.ReadFile("inputs/day2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Println("Contents of file:", string(data))
	temp := strings.Split(string(data), "\n")
	r := regexp.MustCompile("^(\\d+)-(\\d+) (\\w): (\\w+)$")

	count := 0
	for _, line := range temp {
		match := r.FindStringSubmatch(line)

		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		letter := match[3]
		password := match[4]

		if isMatch(min, max, letter[0], password) {
			fmt.Println(line)
			count++
		}
	}
	fmt.Printf("It has %d\n", count)
}

func isMatch(min int, max int, letter byte, password string) bool {
	//count := 0
	X := (password[min-1] == letter)
	Y := (password[max-1] == letter)
	return (X || Y) && !(X && Y)
	/*for i := 0; i < len(password); i++ {

		if password[i] == letter {
			count++
		}
	}
	return count >= min && count <= max*/

}
