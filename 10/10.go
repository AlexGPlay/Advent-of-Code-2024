package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Utils
func readFile() ([][]int, [][]int){
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]int
	var startPoints [][]int

	i:=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		var lineSlice []int

		for j, c := range line{
			value, _ := strconv.Atoi(c)
			lineSlice = append(lineSlice, value)

			if value == 0{
				startPoints = append(startPoints, []int{i, j})
			}
		}

		data = append(data, lineSlice)
		i++
	}

	return data, startPoints
}

func buildKey(position []int) string{
	return strconv.Itoa(position[0]) + "," + strconv.Itoa(position[1])
}

func isValidPosition(data [][]int, position []int, newPosition []int, visited map[string]bool) bool{
	if newPosition[0] < 0 || newPosition[0] >= len(data){
		return false
	}

	if newPosition[1] < 0 || newPosition[1] >= len(data[0]){
		return false
	}

	if visited[buildKey(newPosition)]{
		return false
	}

	currentPositionValue := data[position[0]][position[1]]
	newPositionValue := data[newPosition[0]][newPosition[1]]

	return newPositionValue == currentPositionValue + 1
}

// Part 1
func calculateScore(data [][]int, position []int, visited map[string]bool) int{
	positionKey := buildKey(position)
	
	currentValue := data[position[0]][position[1]]
	visited[positionKey] = true
	if currentValue == 9 {
		return 1
	}
	
	score := 0
	upperPosition := []int{position[0]-1, position[1]}
	if isValidPosition(data, position, upperPosition, visited){
		score += calculateScore(data, upperPosition, visited)
	}

	rightPosition := []int{position[0], position[1]+1}
	if isValidPosition(data, position, rightPosition, visited){
		score += calculateScore(data, rightPosition, visited)
	}

	lowerPosition := []int{position[0]+1, position[1]}
	if isValidPosition(data, position, lowerPosition, visited){
		score += calculateScore(data, lowerPosition, visited)
	}

	leftPosition := []int{position[0], position[1]-1}
	if isValidPosition(data, position, leftPosition, visited){
		score += calculateScore(data, leftPosition, visited)
	}

	return score
}

func part1(){
	data, startPoints := readFile()

	score := 0

	for _, start := range startPoints{
		pathScore := calculateScore(data, start, make(map[string]bool))
		score += pathScore
	}
	
	log.Println(score)
}

// Part 2
func cloneVisited(visited map[string]bool) map[string]bool{
	newVisited := make(map[string]bool)
	for key, value := range visited{
		newVisited[key] = value
	}

	return newVisited
}

func calculateTotalPaths(data [][]int, position []int, visited map[string]bool) int{
	positionKey := buildKey(position)
	newVisited := cloneVisited(visited)
	
	currentValue := data[position[0]][position[1]]
	newVisited[positionKey] = true

	if currentValue == 9 {
		return 1
	}
	
	score := 0
	upperPosition := []int{position[0]-1, position[1]}
	if isValidPosition(data, position, upperPosition, visited){
		score += calculateTotalPaths(data, upperPosition, visited)
	}

	rightPosition := []int{position[0], position[1]+1}
	if isValidPosition(data, position, rightPosition, visited){
		score += calculateTotalPaths(data, rightPosition, visited)
	}

	lowerPosition := []int{position[0]+1, position[1]}
	if isValidPosition(data, position, lowerPosition, visited){
		score += calculateTotalPaths(data, lowerPosition, visited)
	}

	leftPosition := []int{position[0], position[1]-1}
	if isValidPosition(data, position, leftPosition, visited){
		score += calculateTotalPaths(data, leftPosition, visited)
	}

	return score
}

func part2(){
	data, startPoints := readFile()

	score := 0

	for _, start := range startPoints{
		pathScore := calculateTotalPaths(data, start, make(map[string]bool))
		score += pathScore
	}
	
	log.Println(score)
}

func main(){
	part1()
	part2()
}