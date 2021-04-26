# Online Match Service
## Init
```bash
go mod init test.io/api/v1
go list # all modules in this folder
go install github.com/go-delve/delve/cmd/dlv@latest # for remote debuging
code main.go
go run main.go
go build main
go build -ldflags="-s -w" main # optimized build
```

## Run Project
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
  -w /app \
  -v ${PWD}:/app \
  -v /home/dev/go:/go \
  -e GIN_MODE=debug \
  -e PORT=3000 \
  golang:1.16.3-alpine go run main.go
```

## Hot reload
```bash
go get github.com/codegangsta/gin
~/go/bin/gin --appPort 3000 --port 3000
```
or 
```
npx nodemon --exec go run *.go --signal SIGTERM
```

## Local debugging
```bash
go run main.go
Launch file Setting (vscode debuger)
```

### Remote debubg [HAS PROBLEM]
```bash
go get github.com/go-delve/delve/cmd/dlv@latest
cd ~/go/src/work or ~/projectFolder
~/go/bin/dlv debug --headless --log -l :2345 --api-version=2
```


## Test
```bash
http :3000/ping
```