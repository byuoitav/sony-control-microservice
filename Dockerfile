FROM golang:1.6.2

RUN mkdir -p /go/src/github.com/byuoitav
ADD . /go/src/github.com/byuoitav/sony-control

WORKDIR /go/src/github.com/byuoitav/sony-control
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/sony-control"]

EXPOSE 8006
