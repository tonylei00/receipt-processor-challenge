package main

import (
	"encoding/json"
	"net/http"
)

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
