package main

import (
	"fmt"
	"time"
)

func main() {
	// file, err := os.Open("data.csv")
	// if err != nil {
	// 	fmt.Println("Error with opening file")
	// 	return
	// }
	// defer func() {
	// 	if err := file.Close(); err != nil {
	// 		fmt.Println("Error with closing file")
	// 		return
	// 	}
	// }()

	// reader := csv.NewReader(file)

	// for {
	// 	record, err := reader.Read()
	// 	if err != nil {
	// 		fmt.Println("Error with reading line")
	// 		break
	// 	}
	// 	fmt.Println(record[4] == "")
	// }
	val, err := time.Parse("1/02/2006 15:04", "3/20/2021 18:21")
	if err != nil {
		fmt.Println("Incorrect time")
	}
	fmt.Println(val)
}
