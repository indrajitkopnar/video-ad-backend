# Start from the official Go image
FROM golang:1.22-alpine


# Set working directory
WORKDIR /app

# Install Git & other dependencies
RUN apk add --no-cache git

# Copy go mod and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the full source code
COPY . .

# Build the app
RUN go build -o main ./cmd/main.go

# Expose app port
EXPOSE 8080

# Run the binary
CMD ["./main"]
