package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"
)

// Utils
type Robot struct {
	position []int
	velocity []int
}

func readFile() []Robot{
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var robots []Robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersAString := regexp.MustCompile(`-?\d+`).FindAllString(line, -1)
		numbers := make([]int, len(numbersAString))
		for i, numberAsString := range numbersAString {
			fmt.Sscanf(numberAsString, "%d", &numbers[i])
		}

		robot := Robot{position: []int{numbers[0], numbers[1]}, velocity: []int{numbers[2], numbers[3]}}
		robots = append(robots, robot)
	}

	return robots
}

func moveRobot(robot Robot, width int, height int){
	robot.position[0] += robot.velocity[0]
	robot.position[0] %= width
	if robot.position[0] < 0 {
		robot.position[0] += width
	}

	robot.position[1] += robot.velocity[1]
	robot.position[1] %= height
	if robot.position[1] < 0 {
		robot.position[1] += height
	}
}

// Part 1
func calculateSafeFactor(robots []Robot, width int, height int) int{
	y := height / 2	
	x := width / 2

	botsPerQuadrant := [4]int{0, 0, 0, 0}

	for _, robot := range robots {
		if robot.position[0] < x && robot.position[1] < y {
			botsPerQuadrant[0]++
		}
		if robot.position[0] > x && robot.position[1] < y {
			botsPerQuadrant[1]++
		}
		if robot.position[0] < x && robot.position[1] > y {
			botsPerQuadrant[2]++
		}
		if robot.position[0] > x && robot.position[1] > y {
			botsPerQuadrant[3]++
		}
	}
	
	return botsPerQuadrant[0] * botsPerQuadrant[1] * botsPerQuadrant[2] * botsPerQuadrant[3]
}

func part1(times int, width int, height int){
	robots := readFile()

	for _, robot := range robots {
		robot.velocity[0] *= times
		robot.velocity[1] *= times
		moveRobot(robot, width, height)
	}

	safeFactor := calculateSafeFactor(robots, width, height)
	fmt.Println(safeFactor)
}

// Part 2
func createImg(robots []Robot, width int, height int) *image.Gray{
	grid := make([][]bool, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]bool, width)
	}

	for _, robot := range robots {
		grid[robot.position[1]][robot.position[0]] = true
	}

	img := image.NewGray(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			grayValue := uint8(0)
			if grid[i][j] {
				grayValue = 255
			}

			img.SetGray(i, j, color.Gray{Y: grayValue})
		}
	}

	return img
}

func part2(width int, height int){
	robots := readFile()

	for i := 0; i < 10000; i++ {
		for _, robot := range robots {
			moveRobot(robot, width, height)
		}
		
		img := createImg(robots, width, height)
		file, err := os.Create(fmt.Sprintf("imgs/%d.png", i))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			panic(err)
		}
	}
}

func main(){
	// Example => part1(100, 11, 7)
	part1(100, 101, 103)
	// Export all the results to a imgs folder and time to look 1 by 1 😅
	part2(101, 103)
}