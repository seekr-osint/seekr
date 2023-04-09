# Use the official Go image as the base image
FROM golang:1.20-alpine3.16 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go binary
RUN go build -o seekr .

# Use a lightweight base image
FROM alpine:3.15

# Copy the Seekr binary from the build container to the final container
COPY --from=build /app/seekr /app/seekr

# Set the working directory inside the container
WORKDIR /app

# Expose the ports that the application will run on
EXPOSE 5050 8080 8081

# Start the Seekr app
CMD ["./seekr"]

