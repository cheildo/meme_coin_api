# Stage 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest

# Install certificates for secure connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Set the startup command
CMD ["./main"]
