FROM golang:1.14.4-alpine3.11

LABEL maintainer="thiagoluiz.dev@gmail.com"

RUN export GOBIN=$GOPATH/bin

WORKDIR /go/src/ze-delivery-api

COPY . ./

RUN go mod download

RUN go build -o bin/main main.go

CMD ["./bin/main"]

EXPOSE 5001/tcp
