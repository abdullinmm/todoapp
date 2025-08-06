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
RUN CGO_ENABLED=0 GOOS=linux go build -o todoapp ./cmd/todoapp

# --- Stage 2: The final minimum image ---
FROM alpine:latest

WORKDIR /app

# Copy binary from Builder Stage
COPY --from=builder /build/todoapp .

# Optional: Copy Migrations if you need to migrate the base in the container
# COPY --from=builder /build/migrations ./migrations

# Example of a variable environment
ENV PORT=8080

# Open the port (optional, for documentation)
EXPOSE 8080

# Start the application
CMD ["./todoapp"]
