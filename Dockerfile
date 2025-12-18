# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod first for better cache
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/server ./cmd/api

# Runtime stage
FROM scratch

# Copy CA certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary
COPY --from=builder /app/server /server

# Cloud Run uses PORT env var
EXPOSE 8080

ENTRYPOINT ["/server"]
