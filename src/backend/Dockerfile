# Start from the official Golang base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go files
COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download

COPY backend/. .

# Build the app
RUN go build -o main .

# Expose port (change according to your app)
EXPOSE 8080

# Command to run the app
CMD ["./main"]

