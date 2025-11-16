package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    // Create a new request to pass to our handler
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the handler with the ResponseRecorder and request
    handler := http.HandlerFunc(helloHandler)
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body is what we expect
    expected := `
        🌳
        🌲
        🌳
        🌲🌲🌲
        🌳🌳🌳🌳🌳
        🌲🌲🌲🌲🌲🌲🌲
        🌳🌳🌳🌳🌳🌳🌳🌳🌳
        🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲
        🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳
        🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲
        🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳🌳
        🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲🌲
    `

    // Normalize the expected output by trimming spaces
    if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
