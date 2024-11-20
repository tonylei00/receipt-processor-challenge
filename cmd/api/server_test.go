package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Expect 28 points from this receipt
var receipt1 = `{
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
		receiptJson, err := json.Marshal(receipt1)
		if err != nil {
			t.Fatal(err)
		}

		requestBody := bytes.NewReader(receiptJson)
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
