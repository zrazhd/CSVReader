package internal

import "time"

type Company struct {
	ID           int
	TimeGMT      time.Time
	Phone        int
	Organization string
	OLF          string
	Rating       float32
	NumberReview int
	Category     string
	Country      string
	CountryCode  string
	State        string
	City         string
	Street       string
	Building     string
}
