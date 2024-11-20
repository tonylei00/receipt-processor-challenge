package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

const jsonContentType = "application/json"

type Server struct {
	store *DB
	http.Handler
}

type PointsResponse struct {
	Points int `json:"points"`
}

type IDResponse struct {
	ID string `json:"id"`
}

func NewServer(store *DB) *Server {
	server := new(Server)

	server.store = store
	server.Handler = server.routes()

	return server
}

func (s *Server) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /receipts/{id}/points", http.HandlerFunc(s.getReceiptPoints))
	router.Handle("POST /receipts/process", http.HandlerFunc(s.processReceipt))

	return router
}

func (s *Server) getReceiptPoints(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	points, ok := s.store.GetReceiptPointsById(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)

	err := json.NewEncoder(w).Encode(PointsResponse{points})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt

	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	currentRules := []RuleFn{
		onePointForEveryAlphanumericCharInRetailerName(),
		fiftyPointsIfTotalIsRoundDollarAmount(),
		twentyFivePointsIfTotalIsMultipleOfAQuarter(),
		fivePointsForEveryTwoItemsOnReceipt(),
		trimmedLenOfItemDescription(),
		sixPointsIfDayInPurchaseDateIsOdd(),
		tenPointsIfTimeOfPurchaseIsBetween2and4PM(),
	}

	id := uuid.New().String()
	points := receipt.points(currentRules...)

	s.store.SetReceiptPointsById(id, points)

	err = json.NewEncoder(w).Encode(IDResponse{id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
