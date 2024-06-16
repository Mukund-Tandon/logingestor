# Stage 1: Build the Go binary
FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root
RUN apk --no-cache add ca-certificates



COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# # Command to run the executable
ENTRYPOINT ["./main"]
