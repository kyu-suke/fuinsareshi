package service

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/kyu-suke/fuinsareshi/consts"
)

type insertionFunc map[int]func(response *http.Response) error

var InsertionFunc insertionFunc = map[int]func(response *http.Response) error{}

func Proxy() {
	director := func(request *http.Request) {
		request.URL.Scheme = consts.Scheme
		request.URL.Host = consts.Port
	}
	modifyResponse := func(response *http.Response) error {

		if f, ok := InsertionFunc[response.StatusCode]; ok {
			f(response)
		}

		return nil
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}
	server := http.Server{
		Addr:    consts.ProxyPort,
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

func addMethod(i int, f func(response *http.Response) error) {
	InsertionFunc[i] = f
}

func (insertionFunc) Statusok(f func(response *http.Response) error) {
	addMethod(http.StatusOK, f)
}

func (insertionFunc) StatusNotFound(f func(response *http.Response) error) {
	addMethod(http.StatusNotFound, f)
}
