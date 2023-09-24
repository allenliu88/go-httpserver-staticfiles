package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/sirupsen/logrus"
)

var (
	port int
	path string
)

func main() {
	flag.StringVar(&path, "path", "", "静态文件目录")
	flag.IntVar(&port, "port", 8080, "HTTP服务器端口")
	flag.Parse()
	logrus.Println("path: " + path)
	logrus.Println("port: " + strconv.Itoa(port))
	addr := fmt.Sprintf(":%d", port)
	handler := http.FileServer(http.Dir(path))

	// 其实只需要如下1个方法调用即可，但该方法直接阻塞，无法输出监听成功日志，可以拆解
	// http.ListenAndServe(addr, handler)

	// 参考：http.ListenAndServe
	srv := &http.Server{Addr: addr, Handler: http.DefaultServeMux}
	http.Handle("/", handler)

	// Gracefully shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		data := <-sigint
		logrus.Printf("Received Signal: " + data.String())
		logrus.Printf("Start to Shutdown...")

		if err := srv.Shutdown(context.Background()); err != nil {
			logrus.Fatalf("HTTP Server Shutdown: [%v]\n", err)
		}
	}()

	// 参考：srv.ListenAndServe()
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("Internal Server Error When Listen: [%v]\n", err)
		return
	}

	logrus.Printf("HTTP Server Listen on %d\n", port)

	// Serve
	err = srv.Serve(ln)

	if err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Internal Server Error When Serve: [%v]\n", err)
	} else {
		logrus.Printf("HTTP Server Successfully Shutdown: [%v]\n", err)
	}
}
