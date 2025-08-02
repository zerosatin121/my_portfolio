# Use official Go image
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the app
RUN go build -o server .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./server"]
