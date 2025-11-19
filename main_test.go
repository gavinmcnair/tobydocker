package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHtmlHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(htmlHandler)

    // Call the handler with the ResponseRecorder and request
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("htmlHandler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the Content-Type header (allow for charset)
    if contentType := rr.Header().Get("Content-Type"); !strings.HasPrefix(contentType, "text/html") {
        t.Errorf("htmlHandler returned wrong content type: got %v want prefix %v",
            contentType, "text/html")
    }
}

func TestJsHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/snowfall.js", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(jsHandler)

    // Call the handler with the ResponseRecorder and request
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("jsHandler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the Content-Type header (allow for charset)
    if contentType := rr.Header().Get("Content-Type"); !strings.HasPrefix(contentType, "application/javascript") {
        t.Errorf("jsHandler returned wrong content type: got %v want prefix %v",
            contentType, "application/javascript")
    }
}
