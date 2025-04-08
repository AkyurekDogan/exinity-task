# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /

# Copy go.mod and go.sum files
COPY go.mod go.sum config.yml ./
# Download the Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o /go-exinity-task ./cmd/api/main.go

# Stage 2: Run the Go application
FROM alpine:latest

# Expose the port the server will run on
EXPOSE 3000

# Copy the binary from the build stage
COPY --from=builder /go-exinity-task /go-exinity-task
COPY config.yml /config.yml
COPY .env .env

# Command to run the application
CMD ["/go-exinity-task"]
