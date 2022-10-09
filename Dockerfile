FROM golang:1.19-alpine

EXPOSE 80 443

ENV APP_TARGET=prod

RUN apk add --no-cache bash

RUN mkdir "wanpei-backend"

WORKDIR "wanpei-backend"

COPY . .

RUN go build -o ./backend

CMD ["./backend"]









