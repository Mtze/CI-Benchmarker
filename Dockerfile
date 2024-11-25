# Use an official Go runtime as a parent image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy files and download dependencies
COPY . . 
RUN go mod download


# Build the Go application
RUN CGO_ENABLED=0 go build -o benchmarker .

# Start a new stage for the minimal runtime container
FROM gcr.io/distroless/static-debian12

# Set the working directory inside the minimal runtime container
WORKDIR /app

# Copy the built binary from the builder container into the minimal runtime container
COPY --from=builder /app . 

# Run your Go application
CMD ["/app/benchmarker"]