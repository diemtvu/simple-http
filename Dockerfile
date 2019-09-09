# Start from the latest golang base image
FROM golang:latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o demo .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./demo"]
