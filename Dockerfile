FROM golang:1.22.2

WORKDIR /app

COPY . .
RUN go build -o main ./cmd/app/main.go

EXPOSE 8080

CMD ["./main"]