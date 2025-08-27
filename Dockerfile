# --- Stage 1: Application build ---
FROM golang:1.24.2-alpine AS builder

# Install the working directory
WORKDIR /build

# Copy only files for modules for caching Dependency Layer
COPY go.mod go.sum ./
RUN go mod download

# Copy the source in the temporary folder
COPY . .

# Collect binary (Main Package - Way ./cmd/todoApp, correct if the structure is different)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags "-s -w" \
    -o todoapp ./cmd/todoapp

# --- Stage 2: The final minimum image ---
FROM alpine:latest

# Add CA-Certificates for HTTPS queries
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Create Non-Rot user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Copy Binar from Builder Stage
COPY --from=builder /build/todoapp .

# Change the owner
RUN chown appuser:appgroup todoapp

# Switch to Non-ROOT user
USER appuser

ENV PORT=8080
EXPOSE 8080

CMD ["./todoapp"]
