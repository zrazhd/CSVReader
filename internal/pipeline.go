package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

func run {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error with opening file")
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error with closing file")
			return
		}
	}()

	reader := csv.NewReader(file)

	for i := 0; i < 10; i++ {
		record, err := reader.Read()
		if err != nil {
			fmt.Println("Error with reading line")
			break
		}
		fmt.Println(record)
	}

}
