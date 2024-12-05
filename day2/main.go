package main

import (
	"adventofcode2024/internal/utils"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

func mustNum(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("unable to convert string to integer", err)
	}
	return n
}

func convertToInt(data [][]string) [][]int {
	result := make([][]int, len(data))

	for i, record := range data {
		row := strings.Fields(record[0])

		result[i] = make([]int, len(row))

		for j, num := range row {
			result[i][j] = mustNum(num)
		}
	}

	return result
}

func isValid(numbers []int) bool {
	for i := range len(numbers) - 1 {
		diff := numbers[i] - numbers[i+1]
		if diff < 0 {
			diff = -diff
		}

		if !(1 <= diff && diff <= 3) {
			return false
		}
	}

	return slices.IsSorted(numbers) || slices.IsSortedFunc(numbers, func(a, b int) int { return b - a })
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("unable to get current working directory", err)
	}

	data := utils.ReadInput(path.Join(dir, "day2/input.csv"))
	numberList := convertToInt(data)

	a := 0
	b := 0

	for _, numbers := range numberList {
		if isValid(numbers) {
			a++
		}

		valid := false
		for i := range numbers {
			clone := slices.Clone(numbers)
			if isValid(append(clone[:i], clone[i+1:]...)) {
				valid = true
			}
		}

		if valid {
			b++
		}
	}

	fmt.Printf("Safe report count: %d\n", a)
	fmt.Printf("Safe report count tolerating one bad level: %d\n", b)
}
