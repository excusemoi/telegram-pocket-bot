FROM golang:1.18.3-alpine3.16 as builder

COPY . /github.com/excusemoi/telegram-pocket-bot/
WORKDIR /github.com/excusemoi/telegram-pocket-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/excusemoi/telegram-pocket-bot/bin/bot .
COPY --from=0 /github.com/excusemoi/telegram-pocket-bot/configs configs/

EXPOSE 80

CMD ["./bot"]