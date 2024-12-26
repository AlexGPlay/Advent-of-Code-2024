package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

var NUMERIC_KEYPAD = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"", "0", "A"},
}

var MOVEMENT_KEYPAD = [][]string{
	{"", "^", "A"},
	{"<", "v", ">"},
}

type Path struct {
	position []int
	currentPath string
	visited map[string]bool
}

var KEYPAD_CACHE = map[string][]string{}
func calculatePathsInKeypad(keypad [][]string, currentPosition string, objective string) []string{
	key := fmt.Sprintf("%s-%s", currentPosition, objective)
	if _, ok := KEYPAD_CACHE[key]; ok {
		return KEYPAD_CACHE[key]
	}

	positionsMap := make(map[string][]int) 
	for i, row := range keypad {
		for j, value := range row {
			if value == "" {
				continue
			}
			positionsMap[value] = []int{i, j}
		}
	}

	currentPositionCoordinates := positionsMap[currentPosition]
	objectiveCoordinates := positionsMap[objective]

	queue := []Path{{position: currentPositionCoordinates, currentPath: "", visited: map[string]bool{currentPosition: true}}}

	var paths []string
	length := 1000

	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]

		if elem.position[0] == objectiveCoordinates[0] && elem.position[1] == objectiveCoordinates[1] {
			fullPath := elem.currentPath + "A"
			currentLength := len(fullPath)
			if currentLength < length {
				paths = []string{fullPath}
				length = currentLength
			} else if currentLength == length {
				paths = append(paths, fullPath)
			}
			continue
		}

		type Movement struct {
			direction string
			value []int
		}

		for _, movement := range []Movement{{ direction: "^", value: []int{-1,0},}, { direction: "v", value: []int{1,0},}, { direction: "<", value: []int{0,-1},}, { direction: ">", value: []int{0,1},}} {
			newPosition := []int{elem.position[0] + movement.value[0], elem.position[1] + movement.value[1]}
			if newPosition[0] < 0 || newPosition[0] >= len(keypad) || newPosition[1] < 0 || newPosition[1] >= len(keypad[0]) {
				continue
			}

			newPositionKey := keypad[newPosition[0]][newPosition[1]]
			if _, ok := elem.visited[newPositionKey]; ok {
				continue
			}

			if _, ok := positionsMap[newPositionKey]; !ok {
				continue
			}

			newVisited := map[string]bool{}
			for k, v := range elem.visited {
				newVisited[k] = v
			}
			newVisited[newPositionKey] = true

			newPath := elem.currentPath + movement.direction
			queue = append(queue, Path{position: newPosition, currentPath: newPath, visited: newVisited})
		}

	}

	KEYPAD_CACHE[key] = paths
	return paths
}

var MOVEMENTS_CACHE = make(map[string]int)
func calculateMovements(from string, to string, depth int) int{
	key := fmt.Sprintf("%s-%s-%d", from, to, depth)
	if _, ok := MOVEMENTS_CACHE[key]; ok {
		return MOVEMENTS_CACHE[key]
	}

	if depth == 0 {
		paths := calculatePathsInKeypad(MOVEMENT_KEYPAD, from, to)
		return len(paths[0])
	}

	paths := calculatePathsInKeypad(MOVEMENT_KEYPAD, from, to)
	total := math.MaxInt
	for _, path := range paths {
		path = "A" + path
		length := 0
		for i:=0; i<len(path) - 1; i++ {
			from := string(path[i])
			to := string(path[i+1])
			length += calculateMovements(from, to, depth - 1)
		}
		if length < total {
			total = length
		}
	}

	MOVEMENTS_CACHE[key] = total
	return total
}

func cartesianProduct(arrays [][]string) []string {
	if len(arrays) == 0 {
		return []string{}
	}

	if len(arrays) == 1 {
		return arrays[0]
	}

	first := arrays[0]
	rest := cartesianProduct(arrays[1:])

	result := []string{}
	for _, f := range first {
		for _, r := range rest {
			result = append(result, f + r)
		}
	}

	return result
}

func calculateInput(input string, depth int) int{
	var paths [][]string

	currentPosition := "A"
	for _, c := range strings.Split(input, "") {
		res := calculatePathsInKeypad(NUMERIC_KEYPAD, currentPosition, c)
		paths = append(paths, res)
		currentPosition = c
	}
	allPaths := cartesianProduct(paths)

	minPath := math.MaxInt
	for _, path := range allPaths {
		total := 0
		path = "A" + path
		for i:=0; i<len(path) - 1; i++ {
			from := string(path[i])
			to := string(path[i+1])
			total += calculateMovements(from, to, depth - 1)
		}
		if total < minPath {
			minPath = total
		}
	}

	return minPath
}

func runPart(depth int){
	data := readFile()

	sum := 0
	for _, line := range data {
		result := calculateInput(line, depth)
		numberAsString := regexp.MustCompile(`\d+`).FindAllString(line, -1)
		number, _ := strconv.Atoi(numberAsString[0])
		sum += result * number
	}
	fmt.Println(sum)
}

func main(){
	runPart(2)
	runPart(25)
}