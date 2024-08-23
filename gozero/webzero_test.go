package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"testing"
)

func TestWebgozero(t *testing.T) {
	server := rest.MustNewServer(rest.RestConf{
		Host: "0.0.0.0",
		Port: 8888,
	})
	defer server.Stop()

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/hello",
		Handler: helloHandler,
	})

	server.Start()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logx.Infof("Received request: %s", r.URL.Path)
	w.Write([]byte("Hello, World!"))
}
