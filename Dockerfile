# Use the official Golang image (alpine for small size, arm64 for Apple Silicon compatibility)
FROM golang:1.24.5
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main ./main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]