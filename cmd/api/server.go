package main

import (
	"encoding/json"
	"net/http"
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
	ID int `json:"id"`
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
		return
	}
}
func (s *Server) processReceipt(w http.ResponseWriter, r *http.Request) {

}
