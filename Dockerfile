# Use the official Golang image with version 1.23.0 based on Alpine Linux for a lightweight container
FROM golang:1.23.0-alpine

# Set the Current Working Directory inside the container to /app
WORKDIR /app

# Install the latest version of the air package from GitHub
RUN go install github.com/air-verse/air@latest

# Copy the contents of the current directory on the host to the working directory in the container
COPY . .

# Clean up and ensure that the go.mod file is up to date with the dependencies
RUN go mod tidy
