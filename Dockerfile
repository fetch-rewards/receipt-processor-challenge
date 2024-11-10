# Build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git

WORKDIR /app

# Copy the go.mod and go.sum files, and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application and output to /app/receipt-processor
RUN go build -o receipt-processor main.go

# Final stage - use a smaller image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/receipt-processor /app/receipt-processor

# Use JSON syntax for ENTRYPOINT
ENTRYPOINT ["/app/receipt-processor"]

# Label and expose the port
LABEL Name=receiptprocessorchallenge Version=0.0.1
EXPOSE 8080
