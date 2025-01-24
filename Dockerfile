FROM golang:1.23-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /app/out/app


FROM debian:bullseye-slim AS deploy
RUN apt update && apt upgrade
COPY --from=builder /app/app ./
CMD ["./app"]