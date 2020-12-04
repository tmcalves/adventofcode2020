package main

import (
	"container/list"
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
	dayFour()
}

func dayFour() {
	data, err := ioutil.ReadFile("inputs/day4.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	temp := strings.Split(string(data), "\n")

	r := regexp.MustCompile("(\\w+:#?\\w+)")
	list := list.New()
	currentValue := map[string]string{}
	list.PushBack(currentValue)

	for lineNum := 0; lineNum < len(temp); lineNum++ {
		line := temp[lineNum]
		if len(line) == 0 {
			aux := map[string]string{}
			currentValue = aux
			list.PushBack(currentValue)
		}
		match := r.FindAllStringSubmatch(line, -1)
		for _, v := range match {
			keyVal := strings.Split(string(v[1]), ":")
			currentValue[keyVal[0]] = keyVal[1]
		}
	}

	count := 0
	countTotal := 0
	for e := list.Front(); e != nil; e = e.Next() {
		aux := e.Value.(map[string]string)
		countTotal++
		if isValidPassport(aux) {
			count++
		}

	}
	fmt.Printf("Total: %d - Count Valid: %d\n", countTotal, count)

}

func isValidPassport(dict map[string]string) bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" /*"cid"*/}

	for _, key := range keys {
		if val, ok := dict[key]; ok {

			switch key {
			case "byr":
				date, err := strconv.Atoi(val)

				if err != nil {
					fmt.Printf("Invalid to convert byr: %s\n", val)
					return false
				}
				if date < 1920 || date > 2002 {
					fmt.Printf("Invalid byr: %s\n", val)
					return false
				}
				break
			case "iyr":
				date, err := strconv.Atoi(val)

				if err != nil {
					fmt.Printf("Invalid to convert iyr: %s", val)
					return false
				}
				if date < 2010 || date > 2020 {
					fmt.Printf("Invalid iyr: %s\n", val)
					return false
				}
				break
			case "eyr":
				date, err := strconv.Atoi(val)

				if err != nil {
					fmt.Printf("Invalid to convert eyr: %s\n", val)

					return false
				}
				if date < 2020 || date > 2030 {
					fmt.Printf("Invalid iyr: %s\n", val)

					return false
				}
				break
			case "hgt":
				r := regexp.MustCompile("^(\\d+)(cm|in)$")
				match := r.FindStringSubmatch(val)
				if len(match) <= 1 {
					fmt.Printf("Invalid to convert hgt : %s\n", val)

					return false
				}
				intVal, _ := strconv.Atoi(match[1])
				measure := match[2]
				if measure == "cm" {
					if intVal < 150 || intVal > 193 {
						fmt.Printf("Invalid hgt: %s\n", val)
						return false
					}
				} else {
					if intVal < 59 || intVal > 76 {
						fmt.Printf("Invalid hgt: %s\n", val)
						return false
					}
				}
				break
			case "hcl":
				matched, _ := regexp.MatchString("^#[0-9a-f]{6}$", val)
				if !matched {
					fmt.Printf("Invalid hcl: %s\n", val)
					return false
				}
				break
			case "ecl":
				matched, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", val)
				if !matched {
					fmt.Printf("Invalid ecl: %s\n", val)
					return false
				}
				break

			case "pid":
				matched, _ := regexp.MatchString("^\\d{9}$", val)
				if !matched {
					fmt.Printf("Invalid pid: %s\n", val)
					return false
				}
				break

			case "cid":
				break
			default:
				fmt.Printf("This should never happen. %s : %s\n", key, val)
				break
			}

		} else {
			return false
		}
	}

	return true
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
	return (X || Y) && !(X && Y) // Why's there no XOR in GO :(
	/*for i := 0; i < len(password); i++ {

		if password[i] == letter {
			count++
		}
	}
	return count >= min && count <= max*/

}
