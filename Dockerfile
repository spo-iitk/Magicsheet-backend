# Stage 1: Build stage (Pulls code from GitHub)
FROM golang:alpine AS builder

# Install git and certificates
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Clone backend repo directly from GitHub
RUN git clone --depth 1 https://github.com/spo-iitk/Magicsheet-backend.git .

# Download dependencies
RUN go mod download

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

# Expose backend port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
