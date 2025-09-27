# ---------- Build stage ----------
FROM golang:1.25-trixie AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /jobRunner ./cmd/job/main.go

# ---------- Runtime stage ----------
FROM debian:trixie-slim
WORKDIR /
COPY --from=build /jobRunner /jobRunner

ENTRYPOINT ["/jobRunner"]
