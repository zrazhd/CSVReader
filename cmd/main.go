package main

import (
	"fmt"

	"github.com/zrazhd/CSVReader/internal"
)

func main() {
	stats, err := internal.Run("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Total Rows: %v\n", stats.TotalRows)
	fmt.Printf("Not Parsed Rows: %v\n", stats.NotParsedRows)
	fmt.Printf("Invalid Rows: %v\n", stats.InvalidRows)
	fmt.Printf("Success Rows: %v\n", stats.SuccessRows)
}
