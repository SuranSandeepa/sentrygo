# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder 

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o sentrygo ./cmd/server/main.go

# Stage 2: Final lightweight image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/sentrygo .
COPY --from=builder /app/web ./web
EXPOSE 8080
CMD ["./sentrygo"]