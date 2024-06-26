# Start from the official golang image
FROM golang:1.22-alpine AS build

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the code into the container
COPY . .

# Build the Go app
RUN go mod download
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set necessary environment variables
ENV PORT=8080

# Copy the binary file from the previous stage
COPY --from=build /app/main .
COPY --from=build /app/config .

# Expose port 8080 to the outside world
EXPOSE $PORT

# Command to run the executable
CMD ["./main"]
