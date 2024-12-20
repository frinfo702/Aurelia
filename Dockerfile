# Base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o server cmd/server/main.go

# Expose port 8080
EXPOSE 8080

# Set the entry point
CMD ["./server"]
