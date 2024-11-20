package main

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

type RuleFn func(r *Receipt) int

func (receipt *Receipt) points(rules ...RuleFn) int {
	totalReceiptPts := 0

	for _, rule := range rules {
		totalReceiptPts += rule(receipt)
	}

	return totalReceiptPts
}
