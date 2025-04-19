FROM golang:1.23.4 as builder
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o ./main cmd/app/main.go

EXPOSE 8080

ENTRYPOINT ["./main"]
