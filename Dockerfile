# syntax=docker/dockerfile:1

FROM golang:1.16.3-buster as debug
ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get -u github.com/derekparker/delve/cmd/dlv@latest
EXPOSE 3000
CMD [ "go", "run", "main.go" ]


FROM golang:1.16.3-buster as development
ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH
ADD . /go/src/app
WORKDIR /go/src/app


FROM golang:1.16 as build
WORKDIR /go/src/github.com/alexellis/href-counter/
# production
# RUN go get -d -v golang.org/x/net/html 
# only development
COPY /home/dev/go /go 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go


FROM alpine:latest as production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/alexellis/href-counter/main .
CMD ["./main"]  