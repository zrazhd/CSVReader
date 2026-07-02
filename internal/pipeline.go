package internal

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type Stats struct {
	TotalRows     int
	NotParsedRows int
	InvalidRows   int
	SuccessRows   int
}

func Run(path string) (Stats, error) {
	var stats Stats
	file, err := os.Open(path)
	if err != nil {
		return stats, fmt.Errorf("Error with opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.ReuseRecord = true
	_, err = reader.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return stats, nil
		}
		return stats, fmt.Errorf("Failed to read line: %w", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return stats, nil
			}
			return stats, fmt.Errorf("Failed to read line: %w", err)
		}

		stats.TotalRows++

		company, err := Parse(record)
		if err != nil {
			stats.NotParsedRows++
			continue
		}

		err = Validate(company)
		if err != nil {
			stats.InvalidRows++
			continue
		}

		Transform(&company)

		stats.SuccessRows++

	}
}
