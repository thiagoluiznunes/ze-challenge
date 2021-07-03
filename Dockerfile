FROM golang:latest

LABEL maintainer="thiagoluiz.dev@gmail.com"

# RUN apt-get update
# RUN apt-get install vim -y
# RUN apt-get install iputils-ping
RUN export GOBIN=$GOPATH/bin

WORKDIR /go/src/ze-delivery-api

COPY . ./

RUN go mod download

RUN go build -o bin/main main.go

CMD ["./bin/main"]

EXPOSE 8000/tcp
