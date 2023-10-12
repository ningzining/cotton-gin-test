FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 10001

CMD ["/app/main"]