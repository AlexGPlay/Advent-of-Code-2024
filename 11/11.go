package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() map[int]int {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbersMap := make(map[int]int)
	for {
		var number int
		_, err := fmt.Fscanf(file, "%d", &number)
		if err != nil {
			break
		}
		
		if _, ok := numbersMap[number]; ok {
			numbersMap[number]++
		} else {
			numbersMap[number] = 1
		}

	}

	return numbersMap
}

func increaseOrCreate(value int, quantity int, numbers map[int]int) map[int]int {
	if _, ok := numbers[value]; ok {
		numbers[value] += quantity
	} else {
		numbers[value] = quantity
	}
	return numbers
}

func iterate(numbers map[int]int) map[int]int {
	newNumbers := make(map[int]int)

	for value, quantity := range numbers {
		if value == 0 {
			newNumbers = increaseOrCreate(1, quantity, newNumbers)
			continue
		}

		numberAsString := strconv.Itoa(value)
		if len(numberAsString) % 2 == 0 {
			firstHalf := numberAsString[:len(numberAsString)/2]
			firstHalfAsInt, _ := strconv.Atoi(firstHalf)
			newNumbers = increaseOrCreate(firstHalfAsInt, quantity, newNumbers)

			secondHalf := numberAsString[len(numberAsString)/2:]
			secondHalfAsInt, _ := strconv.Atoi(secondHalf)
			newNumbers = increaseOrCreate(secondHalfAsInt, quantity, newNumbers)
			continue
		}

		newNumbers = increaseOrCreate(value * 2024, quantity, newNumbers)
	}
	return newNumbers
}

func iterateTimes(times int){
	numbers := readFile()

	for i := 0; i < times; i++ {
		numbers = iterate(numbers)
	}

	total := 0
	for _, quantity := range numbers {
		total += quantity
	}

	fmt.Println(total)
}

func main() {
	// Part 1
	iterateTimes(25)
	// Part 2
	iterateTimes(75)
}