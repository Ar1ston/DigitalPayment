FROM golang:1.15-rc-alpine
WORKDIR /DigitalPayment
COPY ./ /DigitalPayment
RUN go mod download
WORKDIR /DigitalPayment/Books/Workers
ENTRYPOINT go run  Listener.go