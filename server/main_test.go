package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"smartmirror.server/router"
)

func TestRootEndpoint(t *testing.T) {
	router := router.SetupRouter([]string{"*"})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Smart mirror server is running!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
