package service

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/kyu-suke/fuinsareshi/consts"
)

func Proxy() {

	director := func(request *http.Request) {
		request.URL.Scheme = consts.Scheme
		request.URL.Host = consts.Port
	}
	rp := &httputil.ReverseProxy{
		Director: director,
	}
	server := http.Server{
		Addr:    consts.ProxyPort,
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}

}
