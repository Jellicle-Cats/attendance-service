# Use the official Golang Alpine image as the build stage
FROM golang:1.19-alpine as builder

# Install git (required for fetching dependencies)
RUN apk update && apk add git

# Set the working directory in the image
WORKDIR /usr/src/app

# Copy the source code to the working directory
COPY . .
# Download all dependencies
RUN go get -v ./...

# Build the Go application
RUN go build -o /app

# Use a minimal Alpine Linux image as the final stage
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /home

# Copy the binary from the builder stage to the final image
COPY --from=builder /app ./app

ENV KAFKA_SERVERS = ""

EXPOSE 3456

# Set the command to run the application
CMD ["./app"]