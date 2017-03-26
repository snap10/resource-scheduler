FROM golang

ADD . /go/src/github.com/snap10/resource-scheduler

WORKDIR /go/src/github.com/snap10/resource-scheduler

RUN go get

ENTRYPOINT go run main.go

EXPOSE 8080