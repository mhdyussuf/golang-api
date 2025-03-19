FROM golang:1.20-alpine

WORKDIR /app

# Copy the source code
COPY . .

# Initialize Go module and download dependencies
RUN go mod init golang-api && \
    go mod tidy

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 9800

# Command to run the application
CMD ["./main"]
