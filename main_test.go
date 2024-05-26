package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTomHandler(t *testing.T) {
	// Test GET method
	t.Run("GET method", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(TomHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := Person{Name: "Tom", Age: 20}
		var result Person
		if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}

		if result != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				result, expected)
		}
	})

	// Test POST method
	t.Run("POST method", func(t *testing.T) {
		newPerson := Person{Name: "Jerry", Age: 25}
		body, err := json.Marshal(newPerson)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(TomHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expectedResponse := "OK"
		if rr.Body.String() != expectedResponse {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedResponse)
		}

		if Tom != newPerson {
			t.Errorf("global Tom variable not updated: got %v want %v",
				Tom, newPerson)
		}
	})

	// Test method not allowed
	t.Run("Method not allowed", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(TomHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}

		expectedResponse := "Method not allowed\n"
		if rr.Body.String() != expectedResponse {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedResponse)
		}
	})
}
