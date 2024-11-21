package main

import (
	"fmt"
	"time"
)

const DateTimeNoSecondsLayout = "2006-01-02 15:04"

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (r *Receipt) points(rules ...RuleFn) int {
	totalReceiptPts := 0

	for _, rule := range rules {
		totalReceiptPts += rule(r)
	}

	return totalReceiptPts
}

func (r *Receipt) dateTime() (time.Time, error) {
	dateTime, err := time.Parse(DateTimeNoSecondsLayout, fmt.Sprintf("%v %v", r.PurchaseDate, r.PurchaseTime))
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}
