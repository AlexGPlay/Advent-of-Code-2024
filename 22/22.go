package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Utils
func readFile() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func mix(a int, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func evolveSecretNumber(number int) int {
	newSecretNumber1 := number * 64
	newSecretNumber1 = prune(mix(number, newSecretNumber1))

	newSecretNumber2 := newSecretNumber1 / 32
	newSecretNumber2 = prune(mix(newSecretNumber1, newSecretNumber2))

	newSecretNumber3 := newSecretNumber2 * 2048
	newSecretNumber3 = prune(mix(newSecretNumber2, newSecretNumber3))

	return newSecretNumber3
}

// Part 1
func evolveSecretNumberNTimes(number int, n int) int {
	for i := 0; i < n; i++ {
		number = evolveSecretNumber(number)
	}
	return number
}

func part1() {
	numbers := readFile()
	total := 0
	for _, number := range numbers {
		secretNumber := evolveSecretNumberNTimes(number, 2000)
		total += secretNumber
	}
	println(total)
}

// Part 2
func getLastDigit(number int) int {
	numAsString := strconv.Itoa(number)
	digit := numAsString[len(numAsString)-1]
	lastDigit, err := strconv.Atoi(string(digit))
	if err != nil {
		panic(err)
	}
	return lastDigit
}

func generateSequences(numbers []int) map[string][]int {
	sequences := make(map[string][]int)

	for _, number := range numbers {
		currentSequence := []string{}
		prevValue := getLastDigit(number)
		visited := make(map[string]bool)

		for i := 0; i < 4; i++ {
			number = evolveSecretNumber(number)
			lastDigit := getLastDigit(number)
			diff := lastDigit - prevValue
			prevValue = lastDigit
			currentSequence = append(currentSequence, strconv.Itoa(diff))
		}

		for i := 4; i < 2000; i++ {
			number = evolveSecretNumber(number)
			lastDigit := getLastDigit(number)
			diff := lastDigit - prevValue
			prevValue = lastDigit
			currentSequence = append(currentSequence[1:], strconv.Itoa(diff))

			key := strings.Join(currentSequence, ",")
			if _, ok := visited[key]; ok {
				continue
			} else {
				visited[key] = true
			}

			if _, ok := sequences[key]; ok {
				sequences[key] = append(sequences[key], lastDigit)
			} else {
				sequences[key] = []int{lastDigit}
			}
		}
	}

	return sequences
}

func calculateBestSequence(sequences map[string][]int) int{
	bestSequenceValue := 0

	for _, values := range sequences {
		total := 0
		for _, value := range values {
			total += value
		}
		if total > bestSequenceValue {
			bestSequenceValue = total
		}
	}

	return bestSequenceValue
}

func part2(){
	numbers := readFile()
	sequences := generateSequences(numbers)
	fmt.Println(calculateBestSequence(sequences))
}

func main() {
	part1()
	part2()
}