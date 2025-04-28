# Use the official Go image to build the application
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/user/main.go

# Use a minimal image for the final stage
FROM alpine:latest

# Install ca-certificates for secure connections (if needed)
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/user-service /user-service

# Expose the gRPC port
EXPOSE 50051

# Command to run the application
CMD ["/user-service"]