# Start with the official Golang image as the build environment
FROM golang:1.20.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .
COPY .env .env

# Build the Go app
RUN go build -o expense .

# Start a new, minimal image for the final build
FROM alpine:latest

# Install ca-certificates for HTTPS (optional, but common)
#RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/expense /usr/local/bin/expense
COPY --from=builder /app/.env /app/.env

# Expose the port on which the app will run (optional)
EXPOSE 8080

# Set the command to run the binary
CMD ["expense"]
