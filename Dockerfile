# Use official Golang image to build the Go binary
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files from the server directory
COPY server/go.mod server/go.sum ./

# Install Go dependencies
RUN go mod tidy

# Copy the rest of the application code from the server directory
COPY server/ ./

# Build the Go app
RUN go build -o golang-ws ./cmd

# Create a minimal image to run the Go app
FROM scratch

# Copy the Go binary from the builder stage
COPY --from=builder /app/golang-ws /golang-ws

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["/golang-ws"]
