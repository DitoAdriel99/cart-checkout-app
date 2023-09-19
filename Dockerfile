# Stage 1: Build the Go application
FROM golang:1.19.0-alpine3.15 as builder

WORKDIR /app

COPY . .

# Build the Go application and name the executable as "app"
RUN go build -o app

# Stage 2: Create the final lightweight image
FROM alpine:latest

WORKDIR /app

# Copy the compiled Go application from the builder stage
COPY --from=builder /app/db/migration /app/db/migration
COPY --from=builder /app/app /app/app

EXPOSE $PORT

CMD ["./app"]