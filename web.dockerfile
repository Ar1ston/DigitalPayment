FROM golang:1.15-rc-alpine
RUN apk add --no-cache tzdata
ENV TZ Europe/Moscow