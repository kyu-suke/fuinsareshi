package service

import (
	"net/http"
	"github.com/elazarl/goproxy"
	"github.com/labstack/gommon/log"
)

func Proxy() {
	verbose := false
	addr := ":8080"
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose
	log.Fatal(http.ListenAndServe(addr, proxy))
}
