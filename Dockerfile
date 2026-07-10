# Stage 1: Build stage
FROM golang:1.22-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server

# Stage 2: Final minimal run image
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/main .

# Copy database migrations
COPY --from=builder /app/migrations ./migrations

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
