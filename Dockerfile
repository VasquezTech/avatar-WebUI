# Start from a Go runtime image
FROM golang:latest

RUN apt install make

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . /app

# Build the Go app
RUN make
WORKDIR /app

# Expose port 8050 to the outside world
EXPOSE 8050

# Command to run the executable
CMD ["/app/go-avatar"]
