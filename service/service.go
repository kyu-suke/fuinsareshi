package service

import (
	"net/http"
	"github.com/elazarl/goproxy"
	//"github.com/labstack/gommon/log"
	"log"
)

func Proxy() {
//	verbose := false
//	addr := ":8080"
//	proxy := goproxy.NewProxyHttpServer()
//	proxy.Verbose = verbose
//	log.Fatal(http.ListenAndServe(addr, proxy))
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	log.Fatal(http.ListenAndServe(":8080", proxy))

}
