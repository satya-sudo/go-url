FROM golang:1.21-alpine AS build
WORKDIR /app

# Install git for go mod tidy if needed
RUN apk add --no-cache git

# Copy go modules first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o /gateway ./main.go

# Small runtime image
FROM alpine:3.18
WORKDIR /
COPY --from=build /gateway /gateway

EXPOSE 8080
ENTRYPOINT ["/gateway"]
