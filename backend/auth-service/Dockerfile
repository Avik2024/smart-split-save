# ===== STAGE 1: Build =====
FROM golang:1.20-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Copy go.mod and go.sum
COPY auth-service/go.mod auth-service/go.sum ./auth-service/

WORKDIR /app/auth-service

RUN go mod download

# Copy auth-service and dependencies
COPY auth-service ./auth-service
COPY ../common ../common
COPY ../logger ../logger

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./cmd

# ===== STAGE 2: Minimal image =====
FROM scratch

COPY --from=builder /auth-service /auth-service

EXPOSE 4000

ENTRYPOINT ["/auth-service"]
