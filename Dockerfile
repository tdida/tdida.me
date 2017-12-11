FROM golang:latest

RUN mkdir -p /go/src/tdida.me
WORKDIR /go/src/tdida.me
COPY . /go/src/tdida.me
RUN go get -u github.com/golang/dep/cmd/dep

EXPOSE 8080

CMD ["go", "run", 'main.go']