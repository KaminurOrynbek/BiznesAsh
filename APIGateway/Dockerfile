# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS build

# Set working directory inside the container
WORKDIR /app

# Copy Go mod and sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app (targeting your main.go)
RUN go build -o apigateway ./cmd/gateway/main.go

# Final runtime container
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=build /app/apigateway .

# Expose the port (change if your gateway uses another)
EXPOSE 8080

# Run the binary
CMD ["./apigateway"]