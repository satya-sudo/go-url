# ---------- Build stage ----------
FROM golang:1.25-trixie AS build
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN go build -o /crudService ./cmd/crud/main.go

# ---------- Runtime stage ----------
FROM debian:trixie-slim
WORKDIR /
COPY --from=build /crudService /crudService

EXPOSE 8082
ENTRYPOINT ["/crudService"]
