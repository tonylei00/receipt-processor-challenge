package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tot 28 points
const receipt1 = `{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}`

func TestProcessReceipt(t *testing.T) {
	testDB := NewDB()
	server := NewServer(testDB)

	var receipt1Id string

	t.Run("server responds with ok and json body with id given a valid receipt", func(t *testing.T) {
		requestBody := bytes.NewReader([]byte(receipt1))
		request := httptest.NewRequest(http.MethodPost, "/receipts/process", requestBody)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertContentType(t, response.Header(), jsonContentType)
		receipt1Id = assertIDResponse(t, response.Body)

	})

	t.Run("receipt1 expected to process with 28 points and record in database", func(t *testing.T) {
		got, ok := server.store.GetReceiptPointsById(receipt1Id)
		if !ok {
			t.Fatal("database failed to record points")
		}
		want := 28

		assertInteger(t, got, want)
	})
}

func TestGetPointsById(t *testing.T) {
	testDB := NewDB()
	testServer := NewServer(testDB)
	testId := "id-101"
	testServer.store.SetReceiptPointsById(testId, 500)

	t.Run("server responds ok with json res body of 500 points when given valid id", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/receipts/%s/points", testId), nil)
		response := httptest.NewRecorder()

		testServer.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertContentType(t, response.Header(), jsonContentType)
		gotPoints := assertPointsResponse(t, response.Body)
		assertInteger(t, gotPoints, 500)
	})
	t.Run("server responds with not found when given invalid id", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/receipts/%s/points", "invalid-id"), nil)
		response := httptest.NewRecorder()

		testServer.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}
