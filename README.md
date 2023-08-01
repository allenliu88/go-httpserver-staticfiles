# Go HTTP Server for Static Files

## Local Run

```shell
go run main.go -path=`$pwd`

INFO[0000] HTTP Server Listen on 8080                   
^C
INFO[0000] Received Signal: interrupt                   
INFO[0000] Start to Shutdown...                         
INFO[0000] HTTP Server Successfully Shutdown: [http: Server closed] 
```

## Build & Run

For Linux

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-httpserver-staticfiles-linux-amd64 main.go

chmod +x go-httpserver-staticfiles-linux-amd64
nohup ./go-httpserver-staticfiles-linux-amd64 -path=`$pwd` > http.log 2>&1 &
```
