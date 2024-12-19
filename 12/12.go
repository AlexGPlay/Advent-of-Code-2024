package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Utils
func readFile() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		data = append(data, line)
	}

	return data
}

func isValidPosition(data [][]string, newPosition []int) bool{
	return newPosition[0] >= 0 && newPosition[0] < len(data) && newPosition[1] >= 0 && newPosition[1] < len(data[0])
}

func buildKey(position []int) string {
	return strconv.Itoa(position[0]) + "," + strconv.Itoa(position[1])
}

func walk(data [][]string, position[] int, value string, key int, visited map[string]bool, sections map[int][][]int) {
	if data[position[0]][position[1]] != value {
		return
	}

	positionKey := buildKey(position)
	if visited[positionKey] {
		return
	}

	visited[positionKey] = true
	sections[key] = append(sections[key], position)

	potentialPosition := [][]int{
		{position[0] + 1, position[1]},
		{position[0] - 1, position[1]},
		{position[0], position[1] + 1},
		{position[0], position[1] - 1},
	}

	for _, newPosition := range potentialPosition {
		if isValidPosition(data, newPosition) {
			walk(data, newPosition, value, key, visited, sections)
		}
	}
}

func buildSections(data [][]string) map[int][][]int {
	visited := make(map[string]bool)
	sections := make(map[int][][]int)
	key := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			position := []int{i, j}
			if visited[buildKey(position)] {
				continue
			}

			value := data[i][j]
			walk(data, position, value, key, visited, sections)
			key++
		}
	}

	return sections
}

// Part 1
func calculatePerimeter(section [][]int) int {
	var perimeter int

	for i := 0; i < len(section); i++ {
		position := section[i]
		connections := 0

		for j := 0; j < len(section); j++ {
			if i == j {
				continue
			}

			diff := math.Abs(float64(position[0] - section[j][0])) + math.Abs(float64(position[1] - section[j][1]))
			if diff == 1 {
				connections++
			}
		}
		if connections <= 3 {
			perimeter += 4 - connections
		}
	}

	return perimeter
}

func part1(){
	data := readFile()
	sections := buildSections(data)

	total := 0
	for _, section := range sections {
		total += calculatePerimeter(section) * len(section)
	}

	fmt.Println(total)
}

// Part 2
type PotentialVertex struct {
	position []int
	canOverlap [][]int
}

func isInSection(section [][]int, position []int) bool {
	for _, sectionPosition := range section {
		if sectionPosition[0] == position[0] && sectionPosition[1] == position[1] {
			return true
		}
	}

	return false
}

func getSectionRanges(section [][]int) (int, int, int, int) {
	var minX, minY, maxX, maxY int

	for _, position := range section {
		if position[0] < minX {
			minX = position[0]
		}
		if position[0] > maxX {
			maxX = position[0]
		}
		if position[1] < minY {
			minY = position[1]
		}
		if position[1] > maxY {
			maxY = position[1]
		}
	}

	return minX - 1, minY - 1, maxX + 1, maxY + 1
}

func calculateVertices(section [][]int) int {
	vertices := 0

	minX, minY, maxX, maxY := getSectionRanges(section)

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			position := []int{i, j}

			if isInSection(section, position) {
				continue
			}

			potentialVertices := []PotentialVertex {
				{position: []int{position[0] + 1, position[1] + 1}, canOverlap: [][]int{{position[0] + 1, position[1]}, {position[0], position[1] + 1}}},
				{position: []int{position[0] + 1, position[1] - 1}, canOverlap: [][]int{{position[0] + 1, position[1]}, {position[0], position[1] - 1}}},
				{position: []int{position[0] - 1, position[1] + 1}, canOverlap: [][]int{{position[0] - 1, position[1]}, {position[0], position[1] + 1}}},
				{position: []int{position[0] - 1, position[1] - 1}, canOverlap: [][]int{{position[0] - 1, position[1]}, {position[0], position[1] - 1}}},
			}

			for _, potentialVertex := range potentialVertices {
				count := 0
				for _, overlap := range potentialVertex.canOverlap {
					if isInSection(section, overlap) {
						count += 1
					}
				}

				if count == 2 {
					vertices += 1
				}

				if count == 0 {
					if isInSection(section, potentialVertex.position) {
						vertices += 1
					}
				}
			}

		}
	}

	return vertices
}

func part2(){
	data := readFile()
	sections := buildSections(data)

	total := 0
	for _, section := range sections {
		total += calculateVertices(section) * len(section)
	}
	fmt.Println(total)
}

func main(){
	part1()
	part2()
}