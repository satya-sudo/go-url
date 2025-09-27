# ---------- Build stage ----------
FROM golang:1.25-trixie AS build
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN go build -o /redirectService ./cmd/redirect/main.go

# ---------- Runtime stage ----------
FROM debian:trixie-slim
WORKDIR /
COPY --from=build /redirectService /redirectService

EXPOSE 8083
ENTRYPOINT ["/redirectService"]
