# Use the official Go image as the base image
FROM golang:1.20-alpine3.16 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

RUN go mod download
RUN apk add --no-cache nodejs npm

# Install TypeScript
RUN npm install -g typescript

RUN go generate ./...

RUN tsc --project web --watch false

# Build the Go binary
RUN go build -o seekr main.go

# Use a lightweight base image
FROM alpine:3.18

# Copy the Seekr binary from the build container to the final container
COPY --from=build /app/seekr /bin/seekr

# Set the working directory inside the container
WORKDIR /app

# Expose the ports that the application will run on
EXPOSE 8569

# Start the Seekr app
CMD ["/bin/seekr","--ip","0.0.0.0","--port","8569"]

