package main

import (
    "log"
    "net/http"
)

const PORT = ":8080" // Define the port constant

// Logging middleware to log requests
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Serving request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// Handler for serving the index.html file
func htmlHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

// Handler for serving the snowfall.js file
func jsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/javascript") 
    http.ServeFile(w, r, "snowfall.js")
}

func main() {
    // Serve the HTML file at the root
    http.HandleFunc("/", htmlHandler)
    // Serve the JavaScript file at /snowfall.js
    http.HandleFunc("/snowfall.js", jsHandler)

    // Wrap the default mux with the logging middleware
    loggedMux := loggingMiddleware(http.DefaultServeMux)

    log.Printf("Server is listening on port %s...", PORT)
    // Log errors if the server fails to start
    if err := http.ListenAndServe(PORT, loggedMux); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
