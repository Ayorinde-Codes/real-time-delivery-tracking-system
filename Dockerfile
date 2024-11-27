# Use Golang base image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy application files
COPY . .

# Build the Go binary
RUN go build -o server ./server/main.go

# Expose application port
EXPOSE 8080

# Start the server
CMD ["./server"]
