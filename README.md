# Go HTTP Server for Static Files

## Run

```shell
go run main.go -path=`$pwd`

INFO[0000] HTTP Server Listen on 8080                   
^C
INFO[0000] Received Signal: interrupt                   
INFO[0000] Start to Shutdown...                         
INFO[0000] HTTP Server Successfully Shutdown: [http: Server closed] 
```

## Build

For Linux

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-httpserver-staticfiles-linux-amd64 main.go
```
