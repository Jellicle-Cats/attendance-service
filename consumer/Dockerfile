# Use an official Go runtime as the base image
FROM golang:1.19

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o main

RUN chmod +x main

# Define the command to run your Go application
CMD ["./main"]
