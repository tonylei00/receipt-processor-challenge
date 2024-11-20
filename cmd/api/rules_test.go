package main

import "testing"

var receipt2 = Receipt{
	Retailer:     "M&M Corner Market",
	PurchaseDate: "2022-03-20",
	PurchaseTime: "14:33",
	Items: []Item{
		{"Gatorade", "2.25"},
		{"Gatorade", "2.25"},
		{"Gatorade", "2.25"},
		{"Gatorade", "2.25"},
	},
	Total: "9.00",
}

func TestRules(t *testing.T) {
	type TestCase struct {
		description string
		fn          RuleFn
		want        int
	}

	tests := []TestCase{
		{
			description: "One point for every alphanumeric character in retailer name",
			fn:          onePointForEveryAlphanumericCharInRetailerName,
			want:        14,
		},
		{
			description: "50 points if the total is a round dollar amount with no cents",
			fn:          fiftyPointsIfTotalIsRoundDollarAmount,
			want:        50,
		},
		{
			description: "25 points if the total is a multiple of 0.25",
			fn:          twentyFivePointsIfTotalIsMultipleOfAQuarter,
			want:        25,
		},
		{
			description: "5 points for every two items on the receipt",
			fn:          fivePointsForEveryTwoItemsOnReceipt,
			want:        10,
		},
		{
			description: "Item price * 0.2 if trimmed len of item desc is a mulitple of 3",
			fn:          trimmedLenOfItemDescription,
			want:        0,
		},
		{
			description: "6 points if the day in the purchase date is odd",
			fn:          sixPointsIfDayInPurchaseDateIsOdd,
			want:        0,
		},
		{
			description: "10 points if the time of purchase is between 2 and 4pm",
			fn:          tenPointsIfTimeOfPurchaseIsBetween2and4PM,
			want:        10,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.fn(&receipt2)
			if got != test.want {
				t.Errorf("got=%d, want=%d", got, test.want)
			}
		})
	}
}
