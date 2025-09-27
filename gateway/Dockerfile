FROM golang:1.25-trixie AS build
WORKDIR /app

# Copy go modules first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o /gateway ./cmd/gateway/main.go

# Small runtime image
FROM debian:trixie-slim
WORKDIR /
COPY --from=build /gateway /gateway

EXPOSE 8080
ENTRYPOINT ["/gateway"]
