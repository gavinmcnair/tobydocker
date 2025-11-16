FROM golang:1.23 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go go.mod .

# Run tests
RUN go test -v .

# Build the Go application
RUN go build -o hellotoby .

# Step 2: Create a minimal scratch container
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/hellotoby /hellotoby

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["/hellotoby"]
