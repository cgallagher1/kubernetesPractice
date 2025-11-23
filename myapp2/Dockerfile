FROM golang:1.21-alpine

WORKDIR /app
COPY main.go .

RUN go build -o myapp2 main.go

EXPOSE 5000
CMD ["./myapp2"]
