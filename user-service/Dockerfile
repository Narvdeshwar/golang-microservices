# Use official Go 1.25 Alpine image
FROM golang:1.25.0-alpine AS builder

# Set working directory
WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go mod tidy
RUN go build -o main .

# Use a minimal image for running the app
FROM alpine:latest
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/main .

# Expose service port
EXPOSE 8081

# Run the binary
CMD ["./main"]
