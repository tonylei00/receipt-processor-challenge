package main

import (
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
