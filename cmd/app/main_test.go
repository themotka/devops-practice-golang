package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddition(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{
		{2, 2, 4},
		{-1, -1, -2},
		{-1, 1, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			res := addition(tt.a, tt.b)
			if res != tt.result {
				t.Errorf("addition(%d, %d) = %d; want %d", tt.a, tt.b, res, tt.result)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	expected := "Sample local server 1"
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
