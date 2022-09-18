FROM golang:1.15-rc-alpine
ENV HOME /root
ENV GOPATH ${HOME}/go
ENV PATH ${PATH}:${GOPATH}/bin:/usr/local/go/bin
WORKDIR ${HOME}

RUN apk add --no-cache git

COPY ./ ${GOPATH}/src/DigitalPayment

RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel

EXPOSE 9000
WORKDIR ${GOPATH}/src/DigitalPayment/Web

RUN go mod download
WORKDIR ${GOPATH}/src/DigitalPayment
ENTRYPOINT revel run Web