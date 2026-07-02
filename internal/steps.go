package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Parse(data []string) (Company, error) {

	if len(data) != 14 {
		return Company{}, fmt.Errorf("Bad columns count, expected 14 columns")
	}

	//ID
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return Company{}, fmt.Errorf("Parse ID failed: %w", err)
	}

	//TimeGMT
	t, err := time.Parse("1/02/2006 15:04", data[1])
	if err != nil {
		return Company{}, fmt.Errorf("Parse Time failed: %w", err)
	}
	//Phone
	ph, err := strconv.Atoi(data[2])
	if err != nil {
		return Company{}, fmt.Errorf("Parse Phone failed: %w", err)
	}

	//Rating
	r, err := strconv.ParseFloat(data[5], 32)
	if err != nil {
		return Company{}, fmt.Errorf("Parse Phone failed: %w", err)
	}

	//NumberReview
	nr, err := strconv.Atoi(data[6])
	if err != nil {
		return Company{}, fmt.Errorf("Parse Number Review failed: %w", err)
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

func Validate(c Company) error {
	//ID
	if c.ID <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}
	//Phone
	if c.Phone <= 0 {
		return fmt.Errorf("Phone must be greater than 0")
	}

	//Organization
	if strings.TrimSpace(c.Organization) == "" {
		return fmt.Errorf("Organization can't empty")
	}

	//Rating
	if c.Rating < 0 || c.Rating > 5 {
		return fmt.Errorf("Incorrect Rating, it must be between 0 and 5")
	}

	//NumberReview
	if c.NumberReview < 0 {
		return fmt.Errorf("Negative Number Review")
	}

	//CountryCode

	if c.CountryCode != "" && len(c.CountryCode) != 2 && len(c.CountryCode) != 3 {
		return fmt.Errorf("Country code must contain 2 or 3 letters")
	}

	return nil
}

func Transform(c *Company) {
	c.CountryCode = strings.ToUpper(c.CountryCode)

	org := []rune(c.Organization)

	if len(org) > 0 && !unicode.IsUpper(org[0]) {
		org[0] = unicode.ToUpper(org[0])
		c.Organization = string(org)
	}
}
