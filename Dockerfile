# Start with the official Go image for building the app
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary (main.go is in the root directory)
RUN go build -o receipt-processor ./main.go

# Start with a minimal image for running the app
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/receipt-processor /app/

# Expose the application port
EXPOSE 3000

# Run the binary
CMD ["./receipt-processor"]