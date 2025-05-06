# Stage 1: Build stage
FROM golang:1.21-alpine AS builder

RUN apk update && apk add --no-cache bash coreutils ncurses git

WORKDIR /app
COPY . . 


RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app

# Stage 2: Final image
FROM alpine:latest

# Set environment variable for terminal color support
ENV TERM=xterm-256color

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app . 


# Set the entrypoint to run the application
ENTRYPOINT ["./app"]
