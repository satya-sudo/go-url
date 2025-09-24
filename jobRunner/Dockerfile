# ---------- Build stage ----------
FROM golang:1.21-alpine AS build
WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /jobRunner ./cmd/job/main.go

# ---------- Runtime stage ----------
FROM alpine:3.18
WORKDIR /
COPY --from=build /jobRunner /jobRunner

ENTRYPOINT ["/jobRunner"]
