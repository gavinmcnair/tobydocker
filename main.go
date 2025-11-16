package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Log the request details
    ip := r.RemoteAddr
    userAgent := r.UserAgent()
    log.Printf("Received request from %s with User-Agent: %s", ip, userAgent)

    // Respond with a fancy tree
    tree := `
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
    fmt.Fprintln(w, tree)
}

func main() {
    http.HandleFunc("/", helloHandler)
    log.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
