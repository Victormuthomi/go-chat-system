# 1️⃣ Build Stage - Uses official Golang image
FROM golang:1.23.2 AS builder

# Set working directory inside container
WORKDIR /app

# Copy Go module files first (caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy && go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o chat-system

# 2️⃣ Final Stage - Minimal production image
FROM alpine:latest  

# Set working directory
WORKDIR /root/

# Create a non-root user
RUN adduser -D appuser
USER appuser

# Copy the built binary from builder
COPY --from=builder /app/chat-system .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./chat-system"]

