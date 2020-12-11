package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	tree      byte = '#'
	openSpace byte = '.'
	stepDown  int  = 2
	stepRight int  = 1

	//day11
	freeSit  = 'L'
	floor    = '.'
	occupied = '#'
)

func main() {
	dayEleven()
}

func dayEleven() {
	lines := readFile("inputs/day11.txt")
	grid := [][]byte{}
	for i := 0; i < len(lines); i++ {
		gridLine := []byte{}
		line := lines[i]
		for x := 0; x < len(line); x++ {
			gridLine = append(gridLine, line[x])
		}
		grid = append(grid, gridLine)

	}
	printGrid(grid)

	originalGrid := [][]byte{}
	originalGrid = grid

	newGrid := switchGrid(grid)
	printGrid(newGrid)
	for !areGridsEqual(grid, newGrid) {
		auxGrid := switchGrid(newGrid)
		printGrid(auxGrid)
		grid = newGrid
		newGrid = auxGrid
	}

	found := false

	grid2 := [][]byte{}
	grid2 = originalGrid
	//copy(grid2, originalGrid)
	fmt.Println("GRID 2")
	printGrid(grid2)

	for !found {
		newGrid2 := switchGridPart2(grid2)
		found = areGridsEqual(grid2, newGrid2)

		grid2 = newGrid2
		printGrid(grid2)

	}

	fmt.Printf("There are %d occupied sits for part 1\n", countGridOccupiedSits(newGrid))
	fmt.Printf("There are %d occupied sits for part 2\n", countGridOccupiedSits(grid2))

}
func countGridOccupiedSits(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for x := 0; x < len(grid[i]); x++ {
			if grid[i][x] == occupied {
				count++
			}
		}
	}
	return count
}
func areGridsEqual(grid [][]byte, grid2 [][]byte) bool {

	for i := 0; i < len(grid); i++ {
		for x := 0; x < len(grid[i]); x++ {
			if grid[i][x] != grid2[i][x] {
				return false
			}
		}
	}

	return true
}

func printGrid(grid [][]byte) {
	fmt.Println("=================")
	for i := 0; i < len(grid); i++ {
		fmt.Println(string(grid[i]))
	}
}

func switchGrid(grid [][]byte) [][]byte {
	newGrid := [][]byte{}
	for i := 0; i < len(grid); i++ {
		gridLine := []byte{}
		for x := 0; x < len(grid[i]); x++ {
			b := getNewChar(grid, i, x)
			gridLine = append(gridLine, b)
		}
		newGrid = append(newGrid, gridLine)
	}

	return newGrid
}

func switchGridPart2(grid [][]byte) [][]byte {
	newGrid := [][]byte{}
	for i := 0; i < len(grid); i++ {
		gridLine := []byte{}
		for x := 0; x < len(grid[i]); x++ {
			b := getNewCharV2(grid, i, x)
			gridLine = append(gridLine, b)
		}
		newGrid = append(newGrid, gridLine)
	}

	return newGrid
}

func getNewChar(grid [][]byte, currentIndexGrid int, currentIndexLine int) byte {
	toReturn := grid[currentIndexGrid][currentIndexLine]
	switch toReturn {
	case freeSit:
		if adjentOccupied(grid, currentIndexGrid, currentIndexLine) == 0 {
			return occupied
		}
		break

	case occupied:
		if adjentOccupied(grid, currentIndexGrid, currentIndexLine) >= 4 {
			return freeSit
		}
		break

	default:
		break
	}

	return toReturn
}

func getNewCharV2(grid [][]byte, currentIndexGrid int, currentIndexLine int) byte {
	toReturn := grid[currentIndexGrid][currentIndexLine]
	switch toReturn {
	case freeSit:
		if areOccupiedVisible(grid, currentIndexGrid, currentIndexLine) == 0 {
			return occupied
		}
		break

	case occupied:
		if areOccupiedVisible(grid, currentIndexGrid, currentIndexLine) >= 5 {
			return freeSit
		}
		break

	default:
		break
	}

	return toReturn
}

func adjentOccupied(grid [][]byte, currentIndexGrid int, currentIndexLine int) int {

	count := 0
	for i := currentIndexGrid - 1; i <= currentIndexGrid+1; i++ {
		if i >= 0 && i < len(grid) {
			for x := currentIndexLine - 1; x <= currentIndexLine+1; x++ {
				if x >= 0 && x < len(grid[i]) && !(x == currentIndexLine && i == currentIndexGrid) {
					if grid[i][x] == occupied {
						count++
					}

				}

			}
		}
	}
	return count
}

func areOccupiedVisibleAllTheWay(grid [][]byte, currentIndexGrid int, currentIndexLine int) int {

	count := 0
	for i := 1; i < len(grid); i++ {
		for x := 1; x < len(grid[i]); x++ {
			count += isOcupiedP1(grid, currentIndexGrid-i, currentIndexLine)   //top
			count += isOcupiedP1(grid, currentIndexGrid+i, currentIndexLine)   //down
			count += isOcupiedP1(grid, currentIndexGrid-i, currentIndexLine-x) //top left
			count += isOcupiedP1(grid, currentIndexGrid-i, currentIndexLine+x) // top right
			count += isOcupiedP1(grid, currentIndexGrid+i, currentIndexLine+x) // down right
			count += isOcupiedP1(grid, currentIndexGrid+i, currentIndexLine-x) // down left
			count += isOcupiedP1(grid, currentIndexGrid, currentIndexLine+x)   // right
			count += isOcupiedP1(grid, currentIndexGrid, currentIndexLine-x)   // left
		}
	}
	return count
}

func areOccupiedVisible(grid [][]byte, currentIndexGrid int, currentIndexLine int) int {

	count := 0
	found := [8]bool{}
	for i := 1; i < len(grid); i++ {
		if !found[0] {
			aux, f := isOcupied(grid, currentIndexGrid-i, currentIndexLine) //top
			count += aux
			if f {
				found[0] = true
			}
		}
		if !found[1] {
			aux, f := isOcupied(grid, currentIndexGrid+i, currentIndexLine) //down
			count += aux
			if f {
				found[1] = true
			}
		}
		if !found[2] {
			aux, f := isOcupied(grid, currentIndexGrid-i, currentIndexLine-i) //top left
			count += aux
			if f {
				found[2] = true
			}
		}
		if !found[3] {
			aux, f := isOcupied(grid, currentIndexGrid-i, currentIndexLine+i) // top right
			count += aux
			if f {
				found[3] = true
			}
		}
		if !found[4] {
			aux, f := isOcupied(grid, currentIndexGrid+i, currentIndexLine+i) // down right
			count += aux
			if f {
				found[4] = true
			}
		}
		if !found[5] {
			aux, f := isOcupied(grid, currentIndexGrid+i, currentIndexLine-i) // down left
			count += aux
			if f {
				found[5] = true
			}
		}
		if !found[6] {
			aux, f := isOcupied(grid, currentIndexGrid, currentIndexLine+i) // right
			count += aux
			if f {
				found[6] = true
			}
		}
		if !found[7] {
			aux, f := isOcupied(grid, currentIndexGrid, currentIndexLine-i) // left
			count += aux
			if f {
				found[7] = true
			}
		}
	}
	return count
}

func isOcupiedP1(grid [][]byte, currentIndexGrid int, currentIndexLine int) int {
	if currentIndexGrid < 0 || currentIndexGrid >= len(grid) || currentIndexLine < 0 || currentIndexLine >= len(grid[currentIndexGrid]) {
		return 0
	}

	if grid[currentIndexGrid][currentIndexLine] == occupied {
		return 1
	}
	return 0
}
func isOcupied(grid [][]byte, currentIndexGrid int, currentIndexLine int) (int, bool) {
	if currentIndexGrid < 0 || currentIndexGrid >= len(grid) || currentIndexLine < 0 || currentIndexLine >= len(grid[currentIndexGrid]) {
		return 0, true
	}

	if grid[currentIndexGrid][currentIndexLine] == occupied {
		return 1, true
	}
	return 0, grid[currentIndexGrid][currentIndexLine] != floor
}

func dayTen() {
	lines := readFile("inputs/day10.txt")

	list := []int{}
	list = append(list, 0)
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		list = append(list, val)
	}

	sort.Ints(list)
	oneInc := 0
	threeInc := 0
	otherInc := 0

	for i := 1; i < len(list); i++ {
		diff := list[i] - list[i-1]
		fmt.Printf("%d - %d =  %d\n", list[i], list[i-1], list[i]-list[i-1])
		switch diff {
		case 1:
			oneInc++
			break
		case 3:
			threeInc++
			break
		default:
			fmt.Printf("Found %d\n", diff)
			otherInc++
			break
		}
	}
	//maxList:= list[len(list)-1] +3
	threeInc++

	fmt.Printf("One inc: %d, Three inc: %d, Other: %d, Total: %d\n", oneInc, threeInc, otherInc, len(list))
	fmt.Printf("Result: %d \n", oneInc*threeInc)

	dayTenPart2(list)

}

func dayTenPart2(list []int) {

	thisMap := map[int]bool{}
	for i := 0; i < len(list); i++ {
		thisMap[list[i]] = true
	}

	fmt.Println(thisMap)
	fmt.Println(list)
	lengths := map[int]int{}
	/*val := recCount2(list[len(list)-1], thisMap, 0, "", lengths)

	fmt.Printf("Res: %d\n", val)
	fmt.Println(lengths)
	val = 1*/
	sums := []int{}

	for i := 0; i < len(list); i++ {
		count := countOptions(list, i)
		total := 0
		if i > 6 {
			for x := i; x > i-count; x-- {
				total += sums[x-1]
			}
		} else {
			if i == 0 {
				total = 0
			} else {
				thisMMap := map[int]bool{}
				for t := 0; t < i; t++ {
					thisMMap[list[t]] = true
				}
				fmt.Println(thisMMap)
				total = recCount2(list[i-1], thisMMap, 0, "", lengths)
			}
		}

		sums = append(sums, total)
	}
	fmt.Println(list)
	fmt.Printf("Result is %d\n", sums)
	fmt.Printf("The result is %d\n", sums[len(sums)-1])
	/*count := countOptions(list, len(list)-1)

	fmt.Printf("Init Diff: %d\n", count)

	for i := len(list) - 2; i > 0; i-- {
		diff := countOptions(list, i)
		fmt.Printf("Diff: %d\n", diff)

		count += diff
		fmt.Printf("Calc: %d\n", count)

	}
	fmt.Printf("Res new alg: %d\n", count)*/
}
func countOptions(list []int, current int) int {
	count := 0
	for i := current - 1; i >= 0; i-- {
		diff := list[current] - list[i]
		fmt.Printf("diff %d - %d\n", list[current], list[i])
		if diff <= 3 {
			count++
		} else {
			break
		}
	}
	return count
}

func recCount2(max int, thisMap map[int]bool, current int, path string, lengths map[int]int) int {
	count := 0
	//fmt.Printf("Current %d\n", current)

	path += "-"
	path += strconv.Itoa(current)
	if current >= max {
		lng := len(strings.Split(string(path), "-")) - 1

		fmt.Printf("Reached the end : %s\n", path)
		fmt.Println(lng)
		if _, ok := lengths[lng]; ok {
			lengths[lng]++
		} else {
			lengths[lng] = 1
		}
		count = 1
	} else {
		onePlus := current
		twoPlus := current
		threePlus := current
		onePlus++
		twoPlus += 2
		threePlus += 3
		if _, ok := thisMap[onePlus]; ok {
			//fmt.Printf("%d Has 1p %d\n", current, onePlus)
			count += recCount2(max, thisMap, onePlus, path, lengths)
		}
		if _, ok := thisMap[twoPlus]; ok {
			//fmt.Printf("%d Has 1p %d\n", current, onePlus)
			count += recCount2(max, thisMap, twoPlus, path, lengths)
		}
		if _, ok := thisMap[threePlus]; ok {
			//fmt.Printf("%d Has 3p %d\n", current, threePlus)
			count += recCount2(max, thisMap, threePlus, path, lengths)
		}
	}
	return count

}

func recCount(max int, thisMap map[int]bool, current int, res map[string]bool, path string) int {
	count := 0
	//fmt.Printf("Current %d\n", current)

	path += "-"
	path += strconv.Itoa(current)
	if current >= max {
		fmt.Println("Reached the end")
		res[path] = true
		count = 1
	} else {
		onePlus := current
		threePlus := current
		onePlus++
		threePlus += 3
		if _, ok := thisMap[onePlus]; ok {
			//fmt.Printf("%d Has 1p %d\n", current, onePlus)
			count += recCount(max, thisMap, onePlus, res, path)
		}
		if _, ok := thisMap[threePlus]; ok {
			//fmt.Printf("%d Has 3p %d\n", current, threePlus)
			count += recCount(max, thisMap, threePlus, res, path)
		}

	}
	return count

}

func dayNine() {
	lines := readFile("inputs/day9.txt")
	arr := []int{}
	sum := 0
	total := 0
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		//fmt.Println(arr)
		if len(arr) == 25 {
			sum = val
			count := hasSum(arr, sum)
			if count == 0 {
				//fmt.Printf("%d does not sum\n", sum)
				break
			}
			total += count
			arr = append(arr[:0], arr[1:]...)
		}
		arr = append(arr, val)
	}

	// part 2
	found := false
	for _, line := range lines {
		if found {
			break
		}
		val, _ := strconv.Atoi(line)
		arr = append(arr, val)
		if len(arr) > 1 {
			result := sumIsEqual(arr, sum)

			//fmt.Println(arr)
			//fmt.Printf("Sum: %d. Result: %d\n", sum, result)
			switch result {
			case 0:
				fmt.Print("Found result: ")
				fmt.Println(arr)
				found = true
				break
			case 1:
				for len(arr) > 1 {
					arr = append(arr[:0], arr[1:]...)
					result := sumIsEqual(arr, sum)
					if result == 1 {
						continue
					}
					if result == 0 {
						found = true
						fmt.Print("Found result: ")
						fmt.Println(arr)
						break
					}
					if result == -1 {
						break
					}
				}
				break
			case -1:
				break
			}
		}
	}

	min, max := findMinAndMax(arr)
	fmt.Println(arr)

	fmt.Printf("Min: %d, Max: %d, Sum: %d\n", min, max, min+max)

}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func sumIsEqual(arr []int, sum int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		count += arr[i]
	}

	if count == sum {
		return 0
	}
	if count < sum {
		return -1
	}
	return 1
}
func hasSum(arr []int, sum int) int {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for x := 0; x < len(arr); x++ {
			if arr[i]+arr[x] == sum {
				count++
			}
		}
	}
	return count
}

func dayEight() {
	lines := readFile("inputs/day8.txt")

	_, ended, path := bruteForceFind(-1, lines)
	if !ended {
		for i := len(path) - 1; i >= 0; i-- {
			count, end, _ := bruteForceFind(path[i], lines)
			if end {
				fmt.Printf("Total count is %d when i switched line %d\n", count, i)
			}
		}
	}

	fmt.Println("Done")

}

func bruteForceFind(switchIndex int, lines []string) (int, bool, []int) {
	beenIn := map[int]bool{}
	count := 0
	path := []int{}

	for i := 0; i < len(lines); {

		if _, ok := beenIn[i]; ok {
			return count, false, path
		}
		beenIn[i] = true

		aux := strings.Split(string(lines[i]), " ")
		val := aux[0]
		amount, _ := strconv.Atoi(aux[1])
		switch val {
		case "nop":
			i++
			break
		case "acc":
			i++
			count += amount
			break
		case "jmp":
			if i != switchIndex {
				path = append(path, i)
				i += amount
			} else {
				i++
			}
			break
		}

	}

	return count, true, path
}

// BagSpace for exercise 7
type BagSpace struct {
	bagType string
	count   int
}

func readFile(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		var aux []string
		return aux
	}

	temp := strings.Split(string(data), "\n")
	return temp
}

func daySeven() {

	data, err := ioutil.ReadFile("inputs/day7.txt")
	r := regexp.MustCompile("([A-Za-z0-9 ]+)( bags contain) (\\d+|no) ([A-Za-z0-9 ]+)( bags?)((,) (\\d+) ([A-Za-z0-9 ]+) bags?){0,}(\\.)")
	r2 := regexp.MustCompile("((,) (\\d+) ([A-Za-z0-9 ]+) bags?)")

	withoutBags := 0
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	temp := strings.Split(string(data), "\n")
	thisMap := map[string][]BagSpace{}
	for _, line := range temp {
		match := r.FindStringSubmatch(line)

		initialKey := match[1]
		if match[3] != "no" {
			count, _ := strconv.Atoi(match[3])
			thisMap[initialKey] = append(thisMap[initialKey], BagSpace{count: count, bagType: match[4]})
			match2 := r2.FindAllStringSubmatch(line, -1)
			for _, v := range match2 {
				vCount, _ := strconv.Atoi(v[3])
				thisMap[initialKey] = append(thisMap[initialKey], BagSpace{count: vCount, bagType: v[4]})
			}

		} else {
			withoutBags++
		}
	}

	count := 0
	bag := "shiny gold"
	fmt.Println(thisMap)
	for val, el := range thisMap {
		//fmt.Println(el)

		fmt.Printf("Trying to find for %s in the bag %s\n", bag, val)
		if recFind(thisMap, el, map[string]bool{}, bag, val) {
			count++
		}
	}

	fmt.Printf("Map is %d, without bags is %d and num lines is %d\n", len(thisMap), withoutBags, len(temp))
	fmt.Printf("You can use %d bags\n", count)

	sum := recSum(thisMap, thisMap[bag], map[string]bool{}, bag, bag)
	fmt.Printf("You can carry %d bags\n", sum)

}

func recSum(thisMap map[string][]BagSpace, list []BagSpace, beenIn map[string]bool, bag string, currentKey string) int {

	sum := 0
	for _, v := range list {
		sum += v.count + v.count*recSum(thisMap, thisMap[v.bagType], beenIn, bag, v.bagType)
	}
	return sum
}

func recFind(thisMap map[string][]BagSpace, list []BagSpace, beenIn map[string]bool, bag string, currentKey string) bool {

	for _, v := range list {

		//fmt.Printf("%s - %d\n", v.bagType, v.count)

		if _, ok := beenIn[v.bagType]; ok {
			continue
		}
		beenIn[v.bagType] = true

		if bag == v.bagType {
			fmt.Println("Bag is valid")
			return true
			/*valid := v.count > 1 || len(list) > 1
			if valid {
				fmt.Println("Bag is valid")
				return true
			}*/
		}
		childOld := recFind(thisMap, thisMap[v.bagType], beenIn, bag, v.bagType)
		if childOld {
			return true
		}
	}
	return false
}

func daySix() {
	data, err := ioutil.ReadFile("inputs/day6.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	temp := strings.Split(string(data), "\n")
	list := list.New()
	sMap := map[byte]int{}
	sMap[1] = 0
	list.PushBack(sMap)

	for _, line := range temp {
		if len(line) == 0 {
			aux := map[byte]int{}
			sMap = aux
			sMap[1] = 0
			list.PushBack(sMap)
			continue
		} else {
			sMap[1]++
			for i := 0; i < len(line); i++ {
				if _, ok := sMap[line[i]]; ok {
					sMap[line[i]]++
				} else {
					sMap[line[i]] = 1
				}
			}
		}
	}

	totalCount := 0
	for e := list.Front(); e != nil; e = e.Next() {
		thisMap := e.Value.(map[byte]int)
		//fmt.Printf("Lenght for is %d and people is %d\n", len(thisMap)-1, thisMap[1])
		count := 0
		for val, el := range thisMap {
			if val == 1 {
				continue
			}
			//fmt.Printf("Val is %s and el is %d\n", string(val), el)
			if el == thisMap[1] {
				count++
			}
		}
		totalCount += count
		fmt.Printf("Adding %d\n", count)
	}

	fmt.Printf("The sum is %d\n", totalCount)
}

func dayFive() {
	data, err := ioutil.ReadFile("inputs/day5.txt")

	maxNum := -1
	pass := ""

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	temp := strings.Split(string(data), "\n")

	list := make([]int, len(temp))
	count := 0
	for _, line := range temp {
		row := getRow(line)
		column := getColumn(line)
		id := row*8 + column
		list[count] = id
		count++

		fmt.Printf("Ticket: %s Row: %d Column: %d \n", line, row, column)

		if maxNum < id {
			maxNum = id
			pass = line
		}
	}

	fmt.Printf("Max ID %d is for ID %s \n", maxNum, pass)

	sort.Ints(list)

	for i := 0; i < len(list)-1; i++ {
		if list[i] == (list[i+1] - 2) {
			fmt.Printf("My ID is %d\n", list[i]+1)
			break
		}
	}
}

func getRow(ticket string) int {
	min := 0
	max := 127
	for i := 0; i < 7; i++ {
		calc := float64(max-min) / 2.0
		if ticket[i] == 'F' {
			max -= int(math.Round(calc))

		} else if ticket[i] == 'B' {
			min += int(math.Round(calc))
		}
		fmt.Printf("%s Current row min: %d Current max %d\n", ticket, min, max)
	}

	if ticket[6] == 'F' {
		return int(min)
	} else if ticket[6] == 'B' {
		return int(max)
	}

	return -1
}

func getColumn(ticket string) int {
	min := 0
	max := 7
	for i := 7; i < 10; i++ {
		calc := float64(max-min) / 2.0
		if ticket[i] == 'L' {
			max -= int(math.Round(calc))
		} else if ticket[i] == 'R' {
			min += int(math.Round(calc))
		}
		//fmt.Printf("%s Current col min: %d Current max %d\n", ticket, min, max)
	}

	if ticket[9] == 'L' {
		return min
	} else if ticket[9] == 'R' {
		return max
	}

	return -1
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
