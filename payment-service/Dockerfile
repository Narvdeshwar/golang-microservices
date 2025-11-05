# Use official Go 1.25 image
FROM golang:1.25.0-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Tidy and build the Go app
RUN go mod tidy
RUN go build -o main .

# Use a smaller final image for running the app
FROM alpine:latest
WORKDIR /root/

# Copy built binary from builder stage
COPY --from=builder /app/main .

# Expose port (if your service listens on one)
EXPOSE 8082

# Command to run the binary
CMD ["./main"]
