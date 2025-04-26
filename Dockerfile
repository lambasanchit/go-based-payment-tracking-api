# Stage 1: Build the Go app
FROM golang:1.23 as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the application
RUN go build -o payment-tracking-api main.go

# Stage 2: Run the Go app in a minimal image
FROM alpine:latest

# Install certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built binary from builder stage
COPY --from=builder /app/payment-tracking-api .

# Expose the correct port
EXPOSE 8081

# Run the executable
CMD ["./payment-tracking-api"]
