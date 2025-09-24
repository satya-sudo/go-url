# ---------- Build stage ----------
FROM golang:1.21-alpine AS build
WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN go build -o /crudService ./cmd/crud/main.go

# ---------- Runtime stage ----------
FROM alpine:3.18
WORKDIR /
COPY --from=build /crudService /crudService

EXPOSE 8082
ENTRYPOINT ["/crudService"]
