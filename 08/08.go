package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Utils
func parseFile() (map[string][][]int, int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make(map[string][][]int)

	scanner := bufio.NewScanner(file)
	i := 0
	width := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		width = len(line)

		for j, c := range line {
			if c == "." {
				continue
			}

			if _, ok := data[c]; !ok {
				data[c] = [][]int{}
			}

			data[c] = append(data[c], []int{i, j})
		}

		i++
	}

	return data, i, width
}

func encodePoint(p []int) string {
	return strconv.Itoa(p[0]) + "," + strconv.Itoa(p[1])
}

func isInBounds(p []int, height int, width int) bool{
	if p[0] < 0 || p[0] >= height {
		return false
	}

	if p[1] < 0 || p[1] >= width {
		return false
	}

	return true
}

// Part 1
func calculateAntinode(p1 []int, p2 []int) []int {
	xDiff := p1[0] - p2[0]
	yDiff := p1[1] - p2[1]

	return []int{p1[0] + xDiff, p1[1] + yDiff}
}

func part1() {
	data, height, width := parseFile()
	dataMap := make(map[string]bool)

	for _, v := range data {
		for i, p1 := range v {
			for j, p2 := range v {
				if i == j {
					continue
				}

				antinode := calculateAntinode(p1, p2)

				if isInBounds(antinode, height, width) {
					dataMap[encodePoint(antinode)] = true
				}
			}
		}
	}

	log.Println(len(dataMap))
}

// Part 2
func calculateAntinodes(p1 []int, p2 []int, height int, width int) [][]int {
	xDiff := p1[0] - p2[0]
	yDiff := p1[1] - p2[1]

	point := []int{p1[0], p1[1]}

	antinodes := [][]int{}
	for {
		point = []int{point[0] + xDiff, point[1] + yDiff}

		if !isInBounds(point, height, width) {
			break
		}

		antinodes = append(antinodes, point)
	}

	return antinodes
}

func part2(){
	data, height, width := parseFile()
	dataMap := make(map[string]bool)

	for _, v := range data {
		for i, p1 := range v {
			for j, p2 := range v {
				if i == j {
					continue
				}

				dataMap[encodePoint(p1)] = true
				antinodes := calculateAntinodes(p1, p2, height, width)

				for _, antinode := range antinodes {
					dataMap[encodePoint(antinode)] = true
				}
			}
		}
	}

	log.Println(len(dataMap))
}

func main() {
	part1()
	part2()
}