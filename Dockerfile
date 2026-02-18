# Stage 1: Build the application in a dedicated build environment
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app, creating a static binary for a linux env
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

# Stage 2: Create a minimal production image
FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 8080
CMD ["/main"]