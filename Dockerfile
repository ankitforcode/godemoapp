FROM golang:1.4

ADD . /go/src/

RUN go get github.com/Sirupsen/logrus/hooks/syslog

RUN go get github.com/Sirupsen/logrus

CMD ["go","run","/go/src/main.go"]
