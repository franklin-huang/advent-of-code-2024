package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadInput(filePathStr string) [][]string {
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
