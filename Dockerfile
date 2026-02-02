# Step 1: Build the Go binary
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Step 2: Create a tiny image to run the binary
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./main"]