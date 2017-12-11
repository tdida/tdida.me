FROM golang:latest

RUN mkdir -p /go/src/tdida.me
WORKDIR /go/src/tdida.me
COPY . /go/src/tdida.me
RUN go get -u github.com/beego/bee
RUN go get -d -v
RUN go install -v

EXPOSE 8081

CMD ["bee", "run"]