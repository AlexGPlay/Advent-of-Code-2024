package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile() ([][]string, [][]string) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	all, _ := io.ReadAll(file)
	blocks := strings.Split(string(all), "\n\n")

	var keys [][]string
	var locks [][]string

	for _, block := range blocks {
		elem := strings.Split(block, "\n")

		if elem[0][0] == '#' {
			locks = append(locks, elem)
		} else {
			keys = append(keys, elem)
		}
	}

	return keys, locks
}

func calculateLockHeights(lock []string) []int {
	var heights []int

	for i := 0; i < len(lock[0]); i++ {
		for height, line := range lock {
			if line[i] == '.' {
				heights = append(heights, height)
				break
			}
		}
	}

	return heights
}

func calculateKeyHeights(key []string) []int {
	var heights []int

	for i := 0; i < len(key[0]); i++ {
		for height, line := range key {
			if line[i] == '#' {
				heights = append(heights, len(key) - height)
				break
			}
		}
	}

	return heights
}

func checkKeysAndLocks(lockHeight []int, keys [][]int, maxHeight int) int {
	total := 0

	for _, key := range keys {
		isValid := true
		for i := 0; i < len(key); i++ {
			sum := key[i] + lockHeight[i]
			if sum > maxHeight {
				isValid = false
				break
			}
		}
		if isValid {
			total++
		}
	}

	return total
}

func run() {
	keys, locks := readFile()
	var lockHeights [][]int
	var keyHeights [][]int

	for _, lock := range locks {
		lockHeights = append(lockHeights, calculateLockHeights(lock))
	}
	for _, key := range keys {
		keyHeights = append(keyHeights, calculateKeyHeights(key))
	}


	total := 0

	for _, lockHeight := range lockHeights {
		total += checkKeysAndLocks(lockHeight, keyHeights, len(keys[0]))
	}

	fmt.Println(total)
}

func main() {
	run()
}