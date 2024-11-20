package main

import (
	"math"
	"strconv"
	"strings"
)

// One point for every alphanumeric character in the retailer name.
func onePointForEveryAlphanumericCharInRetailerName(r *Receipt) int {
	isAlphaNumeric := func(c rune) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	}

	count := 0

	for _, char := range r.Retailer {
		if isAlphaNumeric(char) {
			count++
		}
	}

	return count
}

func fiftyPointsIfTotalIsRoundDollarAmount(r *Receipt) int {
	// 50 points if the total is a round dollar amount with no cents.
	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return 0
	}

	if total == math.Trunc(total) {
		return 50
	}

	return 0
}

// 25 points if the total is a multiple of 0.25
func twentyFivePointsIfTotalIsMultipleOfAQuarter(r *Receipt) int {
	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return 0
	}

	if math.Mod(total, 0.25) == 0 {
		return 25
	}

	return 0
}

// 5 points for every two items on the receipt.
func fivePointsForEveryTwoItemsOnReceipt(r *Receipt) int {
	lenItems := len(r.Items)

	if lenItems >= 2 {
		mulitpler := int(math.Floor(float64(lenItems) / 2))
		return 5 * mulitpler
	}

	return 0
}

// If the trimmed length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up to the nearest integer.
// The result is the number of points earned.
func trimmedLenOfItemDescription(r *Receipt) int {
	points := 0

	for _, item := range r.Items {
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))

		if trimmedLen%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}

			points += int(math.Ceil(price * 0.2))
		}
	}

	return points
}

// 6 points if the day in the purchase date is odd
func sixPointsIfDayInPurchaseDateIsOdd(r *Receipt) int {
	isOdd := func(n int) bool {
		return n%2 != 0
	}

	date, err := r.dateTime()
	if err != nil {
		return 0
	}

	if isOdd(date.Day()) {
		return 6
	}

	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func tenPointsIfTimeOfPurchaseIsBetween2and4PM(r *Receipt) int {
	date, err := r.dateTime()
	if err != nil {
		return 0
	}

	if date.Hour() >= 14 && date.Hour() <= 16 {
		return 10
	}

	return 0
}
