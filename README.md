
# Start
go mod init test.io/api/v1
go list # all modules in this folder
go install github.com/go-delve/delve/cmd/dlv@latest # for remote debuging
code main.go
go run main.go
go build main
go build -ldflags="-s -w" main # optimized build

# Running
docker run --rm -p 8080:8080 -p 2345:2345 -w /app -it -v ${PWD}:/app -v /home/dev/go:/go -e GIN_MODE=debug golang:1.16.3-alpine go run main.go
run main.go

## Remote debubg
go get github.com/go-delve/delve/cmd/dlv@latest
cd ~/go/src/work or ~/projectFolder
~/go/bin/dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2

## Docker compose

docker-compose up --force-recreate

## Test
http :8080/ping