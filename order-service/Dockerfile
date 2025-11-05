# Use official Go 1.25 Alpine image
FROM golang:1.25.0-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy dependency files first (for better layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your code
COPY . .

# Run tidy and build
RUN go mod tidy
RUN go build -o main .

# Use a small image to run the final binary
FROM alpine:latest
WORKDIR /root/

# Copy built binary from builder stage
COPY --from=builder /app/main .

# Expose your service port (adjust as needed)
EXPOSE 8083

# Run the app
CMD ["./main"]
