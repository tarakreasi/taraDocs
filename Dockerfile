# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies for Go and Frontend
RUN apk add --no-cache gcc musl-dev nodejs npm

# Copy Go manifests
COPY go.mod go.sum ./
RUN go mod download

# Copy NPM manifests
COPY package.json package-lock.json ./
RUN npm ci

# Copy source code
COPY . .

# Build frontend assets
RUN npm run build

# Build binary
RUN go build -ldflags="-s -w" -o server cmd/server/main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

# Copy binaries
COPY --from=builder /app/server .

# Copy resources (templates, assets)
COPY --from=builder /app/views ./views
COPY --from=builder /app/public ./public

# Prepare docs volume directory
RUN mkdir -p /app/docs

# Expose port
EXPOSE 3000

# Environment variables
ENV PORT=3000
ENV DOCS_PATH=/app/docs

# Run server
CMD ["./server"]
