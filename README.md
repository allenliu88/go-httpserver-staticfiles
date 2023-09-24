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
nohup ./go-httpserver-staticfiles-linux-amd64 -path=`$pwd` -port=8080 > http.log 2>&1 &
```

For Windows

```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/go-httpserver-staticfiles-windows-amd64.exe main.go

chmod +x go-httpserver-staticfiles-windows-amd64.exe
nohup ./go-httpserver-staticfiles-windows-amd64.exe -path=`$pwd` -port=8080 > http.log 2>&1 &
```

## License

Copyright &copy; 2023 Allen

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
