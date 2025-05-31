# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS build

# Set working directory inside the container
WORKDIR /app

# Copy Go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app
RUN go build -o contentservice ./cmd/main.go

# Final runtime container
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=build /app/contentservice .

# Expose the port
EXPOSE 8082

# Run the binary
CMD ["./contentservice"] 