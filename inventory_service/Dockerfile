# Build stage
FROM golang:1.24-alpine3.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o /main ./cmd/httpserver/main.go
RUN go build -o /migrate ./cmd/migrate/main.go

# Run stage
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /main .
COPY --from=builder /migrate .

COPY .env .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]