# Start from a Go runtime image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . /app

# Build the Go app
RUN go mod tidy && go build -o /app/go-avatar
RUN chmod +x go-avatar

# Expose port 8080 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["/app/go-avatar"]
