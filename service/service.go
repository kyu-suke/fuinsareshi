package service

import (
	//"github.com/elazarl/goproxy"
	//"github.com/labstack/gommon/log"

	"log"
	"net/http"
	"net/http/httputil"

	"io/ioutil"
	"fmt"
)

func Proxy() {
//	verbose := false
//	addr := ":8080"
//	proxy := goproxy.NewProxyHttpServer()
//	proxy.Verbose = verbose
//	log.Fatal(http.ListenAndServe(addr, proxy))
//	proxy := goproxy.NewProxyHttpServer()
//	proxy.Verbose = true
//	proxy.NewConnectDialToProxy()
//	log.Fatal(http.ListenAndServe(":8080", proxy))

	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":80"
	}
	modifyResponse := func(response *http.Response) error {
		r, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(r))
		return nil
	}
	rp := &httputil.ReverseProxy{
		Director: director,
		ModifyResponse: modifyResponse,
	}
	server := http.Server{
		Addr:    ":9000",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}

}
