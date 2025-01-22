FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-rest-search-service ./cmd/

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/go-rest-search-service .
COPY config.yaml /root/config.yaml
COPY input.txt /root/input.txt

RUN chmod +x /root/go-rest-search-service

EXPOSE 8080

CMD ["./go-rest-search-service"]
