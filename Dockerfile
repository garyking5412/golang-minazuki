FROM golang:1.24 AS builder
LABEL authors="takeshitran"
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .
FROM debian:bookworm-slim
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE "3004"
CMD ["./main"]