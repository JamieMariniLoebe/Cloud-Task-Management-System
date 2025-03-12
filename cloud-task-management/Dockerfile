FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o cloud-task-management ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cloud-task-management .

CMD ["./cloud-task-management"]