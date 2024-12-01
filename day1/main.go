package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

func readInput(filePathStr string) [][]string {
	f, err := os.Open(filePathStr)
	if err != nil {
		log.Fatal("unable to open file", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("unable to read csv", err)
	}

	return records
}

func parseData(data [][]string) (left []int, right []int) {
	left = make([]int, len(data))
	right = make([]int, len(data))

	for i, record := range data {
		row := strings.Fields(record[0])

		if len(row) != 2 {
			log.Fatal("invalid input data, expected 2 columns")
		}

		num1, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal("unable to convert string to integer", err)
		}

		num2, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal("unable to convert string to integer", err)
		}

		left[i] = num1
		right[i] = num2
	}

	return left, right
}

func calculateDistance(left []int, right []int) int {
	distance := 0

	for i := 0; i < len(left); i++ {
		d := left[i] - right[i]

		if d < 0 {
			distance += -d
		} else {
			distance += d
		}
	}

	return distance
}

func calculateSimilarity(left []int, right []int) int {
	score := 0
	rightListOccurrence := make(map[int]int)

	for _, num := range right {
		if _, ok := rightListOccurrence[num]; ok {
			rightListOccurrence[num]++
		} else {
			rightListOccurrence[num] = 1
		}
	}

	for _, num := range left {
		if count, ok := rightListOccurrence[num]; ok {
			score += num * count
		}
	}

	return score
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("unable to get current working directory", err)
	}

	data := readInput(path.Join(dir, "day1/input.csv"))
	left, right := parseData(data)
	slices.Sort(left)
	slices.Sort(right)

	totalDistance := calculateDistance(left, right)
	similarityScore := calculateSimilarity(left, right)

	fmt.Printf(`
Total distance: %d
Similarity score: %d
`,
		totalDistance,
		similarityScore)

}
