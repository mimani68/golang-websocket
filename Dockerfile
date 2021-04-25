FROM golang:1.16.3-buster

ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH

ADD . /go/src/app
WORKDIR /go/src/app

# RUN go get -u github.com/derekparker/delve/cmd/dlv@latest

EXPOSE 8080

# RUN go build -o main main.go
# CMD ["./main"]

CMD [ "go", "run", "main.go" ]
