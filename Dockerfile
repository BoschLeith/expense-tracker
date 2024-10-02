# Use the official Go image based on Alpine for a lightweight container
FROM golang:1.23.0-alpine

# Install build dependencies for CGO
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Install the latest version of the air package
RUN go install github.com/air-verse/air@latest

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the module dependencies specified in go.mod and go.sum
RUN go mod download
