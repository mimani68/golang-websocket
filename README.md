# Online Match Service

## Technical spec

| Title | Value | Version |
|:---:|:---:|:---:|
| Language | GoLang | v1.16.3 |
| Realtime  | WebSocket | * |
| framework | GOSF http://gosf.io/#client-broadcasts | * |
| Hot reload | nodemon dlv | * |
| development log | Logz.io | * |



## Init
```bash
go mod init blackoak.cloud/balout/v2
go list     # all modules in this folder
go get .    # get all dependencies
go install github.com/go-delve/delve/cmd/dlv@latest # for remote debuging
go get -u github.com/go-delve/delve/cmd/dlv
go get -u github.com/gin-gonic/gin
go get github.com/go-redis/redis/v8
code main.go
go run main.go
```
## Build

> https://docs.docker.com/develop/develop-images/multistage-build/
### normal

```bash
go build main.go

```

### optimized

```bash
go build -ldflags "-w" main.go     # optimized build
go build -ldflags "-s -w" main.go  # best optimized setting
```
### optimized for docker alpine base image

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go
```
then access
```bash
docker run --rm -p 3000:3000 -it -v ${PWD}:/app -w /app alpine:latest ./main
```

## Docker Compose
```bash
docker-compose -p game up --force-recreate
docker-compose -p game down --rm local --remove-orphans
```

## Running in Standalone docker
```bash
docker run \
  --rm \
  -it \
  -p 3000:3000 \
  -p 2345:2345 \
  -w /go/src/app \
  -v ${PWD}:/go/src/app \
  -v /home/dev/go:/go \
  -e GIN_MODE=debug \
  -e PORT=3000 \
  golang:1.16.3-alpine go run main.go
```

## Hot reload
```bash
go get github.com/codegangsta/gin
vim ~/config/server.json
~/go/bin/gin --appPort 3000 --port 3000
```
or for local system that has nodemon
```
nodemon --exec go run main.go --signal SIGTERM
```
and finaly live reload in docker-compose
```bash
nodemon -e go --exec docker-compose -p game up  --force-recreate
```


## Local debugging
```bash
go run main.go
Launch file Setting (vscode debuger)
```

### Remote debubg
```bash
go get github.com/go-delve/delve/cmd/dlv@latest
cd ~/go/src/work or ~/projectFolder
~/go/bin/dlv debug --headless --log -l :2345 --api-version=2
```
or 
```bash
dlv debug -l 0.0.0.0:2345 --headless --log --api-version=2 --accept-multiclient main.go
```

then run 
```bash
ws :3000/balout/api/v1/match/
```
or 
```bash
chrome ${workspace}/client-simulator/index.html
```


## Test
```bash
http :3000/ping
```