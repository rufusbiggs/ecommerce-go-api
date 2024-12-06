# First stage: Build the Go binary
FROM golang:1.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container to download dependencies
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Second stage: Use a minimal runtime image
FROM debian:buster-slim

# Install required certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory in the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary has execution permissions
RUN chmod +x /app/main

# Expose the application port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/app/main"]



