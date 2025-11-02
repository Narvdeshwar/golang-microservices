# Use official Go image
FROM golang:1.22

# Set working directory
WORKDIR /golang-microservices/payment-services

# Copy code into container
COPY . .

# Build the Go app
RUN go mod tidy
RUN go build -o main .

# Run the binary
CMD ["./main"]
