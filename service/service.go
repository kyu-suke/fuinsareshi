package service

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func Proxy() {

	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":8080"
	}
	rp := &httputil.ReverseProxy{
		Director: director,
	}
	server := http.Server{
		Addr:    ":9000",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}

}
