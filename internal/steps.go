package internal

import (
	"fmt"
	"strconv"
	"time"
)

func Parse(data []string) (Company, error) {

	if len(data) < 14 {
		return Company{}, fmt.Errorf("Bad columns count")
	}

	//ID
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return Company{}, fmt.Errorf("Something wrong with ID (It's not a number)")
	}

	//TimeGMT
	t, err := time.Parse("1/02/2006 15:04", data[1])
	if err != nil {
		return Company{}, fmt.Errorf("Something wrong with Time (It's not a time)")
	}
	//Phone
	ph, err := strconv.Atoi(data[2])
	if err != nil {
		return Company{}, fmt.Errorf("Something wrong with Phone (It's not a number)")
	}

	//Rating
	r, err := strconv.ParseFloat(data[5], 32)
	if err != nil {
		return Company{}, fmt.Errorf("Something wrong with Phone (It's not a number)")
	}

	//NumberReview
	nr, err := strconv.Atoi(data[2])
	if err != nil {
		return Company{}, fmt.Errorf("Something wrong with Number Review (It's not a number)")
	}

	return Company{
		ID:           id,
		TimeGMT:      t,
		Phone:        ph,
		Organization: data[3],
		OLF:          data[4],
		Rating:       float32(r),
		NumberReview: nr,
		Category:     data[7],
		Country:      data[8],
		CountryCode:  data[9],
		State:        data[10],
		City:         data[11],
		Street:       data[12],
		Building:     data[13],
	}, nil
}

func Validate(c *Company) {

}

func Transform(c *Company) {

}
