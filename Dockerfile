
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o dsp ./cmd/httpserver_api


FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/dsp .

EXPOSE 8080
ENTRYPOINT ["./dsp"]
