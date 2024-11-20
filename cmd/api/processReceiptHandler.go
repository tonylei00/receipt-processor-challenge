package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt

	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	currentRules := []RuleFn{
		onePointForEveryAlphanumericCharInRetailerName,
		fiftyPointsIfTotalIsRoundDollarAmount,
		twentyFivePointsIfTotalIsMultipleOfAQuarter,
		fivePointsForEveryTwoItemsOnReceipt,
		trimmedLenOfItemDescription,
		sixPointsIfDayInPurchaseDateIsOdd,
		tenPointsIfTimeOfPurchaseIsBetween2and4PM,
	}

	id := uuid.New().String()
	points := receipt.points(currentRules...)

	s.store.SetReceiptPointsById(id, points)

	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(IDResponse{id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
