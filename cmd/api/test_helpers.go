package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d status code, want %d", got, want)
	}
}

func assertContentType(t testing.TB, headers http.Header, want string) {
	t.Helper()

	got := headers.Get("Content-Type")
	if got != want {
		t.Errorf("got %q, want %q content-type", got, want)
	}
}

func assertIDResponse(t testing.TB, got io.Reader) string {
	t.Helper()

	var response IDResponse
	err := json.NewDecoder(got).Decode(&response)
	if err != nil {
		t.Error(err)
	}

	return response.ID
}

func assertPointsResponse(t testing.TB, got io.Reader) int {
	t.Helper()

	var pointsResponse PointsResponse
	err := json.NewDecoder(got).Decode(&pointsResponse)
	if err != nil {
		t.Error(err)
	}

	return pointsResponse.Points
}

func assertInteger(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}
}
