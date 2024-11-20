package main

func onePointForEveryAlphanumericCharInRetailerName() RuleFn {
	// One point for every alphanumeric character in the retailer name.

	isAlphaNumeric := func(c rune) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	}

	return func(r *Receipt) int {
		count := 0

		for _, char := range r.Retailer {
			if isAlphaNumeric(char) {
				count++
			}
		}

		return count
	}
}

func fiftyPointsIfTotalIsRoundDollarAmount() RuleFn {
	// 50 points if the total is a round dollar amount with no cents.
	return func(r *Receipt) int {

	}
}
