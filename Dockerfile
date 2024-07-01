# Use an official Go runtime as the base image
FROM golang:latest  as builder

# Set the working directory in the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose a port that the application will run on
EXPOSE 8000

# Define the command to start the application
CMD ["./myapp"]