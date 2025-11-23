FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -o myapp1 main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp1 .
EXPOSE 5000
CMD ["./myapp1"]
