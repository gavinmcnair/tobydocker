package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Log the request details
    ip := r.RemoteAddr
    userAgent := r.UserAgent()
    log.Printf("Received request from %s with User-Agent: %s", ip, userAgent)

    // Set the content type to HTML
    w.Header().Set("Content-Type", "text/html")

    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // Define tree dimensions
    const treeWidth = 400
    const treeHeight = 600
    const trunkWidth = 40
    const trunkHeight = 60

    // Calculate trunk position
    trunkX := (treeWidth / 2) - (trunkWidth / 2)
    trunkY := treeHeight - trunkHeight

    // Generate random positions for snowflakes
    snowflakes := []string{}
    for i := 0; i < 50; i++ {
        // Random x position for snowflakes within the SVG area
        x := rand.Intn(treeWidth)
        // Random y position for snowflakes within the SVG area
        y := rand.Intn(treeHeight)
        // Random size for snowflakes (between 2 and 10 pixels)
        size := rand.Intn(9) + 2
        snowflakes = append(snowflakes, fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d" fill="white" />`, x, y, size))
    }

    // Respond with an HTML document containing the SVG Christmas tree and snowflakes
    html := fmt.Sprintf(`<!DOCTYPE html>
    <html>
    <head>
        <meta http-equiv="refresh" content="1">
        <title>Christmas Tree with Snow</title>
    </head>
    <body>
    <svg width="%d" height="%d" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg">
        <rect x="%d" y="%d" width="%d" height="%d" fill="saddlebrown" />
        <polygon points="%d,%d %d,%d %d,%d" fill="green" />
        <polygon points="%d,%d %d,%d %d,%d" fill="green" />
        <polygon points="%d,%d %d,%d %d,%d" fill="green" />
        %s
    </svg>
    </body>
    </html>`, 
        treeWidth, treeHeight, treeWidth, treeHeight,
        trunkX, trunkY, trunkWidth, trunkHeight,
        treeWidth/2, 0, 0, treeHeight-200, treeWidth, treeHeight-200,
        treeWidth/2, 100, 50, treeHeight-100, treeWidth-50, treeHeight-100,
        treeWidth/2, 200, 70, treeHeight-50, treeWidth-70, treeHeight-50,
        fmt.Sprintf("%s", snowflakes))

    fmt.Fprintln(w, html)
}

func main() {
    http.HandleFunc("/", helloHandler)
    log.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
