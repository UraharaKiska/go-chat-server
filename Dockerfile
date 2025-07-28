FROM golang:1.23.1-alpine AS builder

COPY . /github.com/UraharaKiska/go-chat-server/source/
WORKDIR /github.com/UraharaKiska/go-chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/UraharaKiska/go-chat-server/source/bin/chat_server .

CMD ["./chat_server"]