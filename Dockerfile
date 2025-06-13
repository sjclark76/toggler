# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
# CGO_ENABLED=0 is used to build a statically linked binary, which is good for minimal base images
# -o /app/toggler specifies the output file name and path
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/toggler ./main.go

# Stage 2: Create the final, minimal image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/toggler .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./toggler"]