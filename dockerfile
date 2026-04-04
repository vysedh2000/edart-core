# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

# Install git (needed by Go modules) and other deps
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ---------- Run Stage ----------
FROM alpine:latest

# Add ca-certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8090
EXPOSE 8090

# Run the binary
CMD ["./main"]
