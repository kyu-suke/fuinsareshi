package service

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/kyu-suke/fuinsareshi/consts"
)

type insertionFuncs map[int]func(response *http.Response) error

var InsertionFuncs insertionFuncs = map[int]func(response *http.Response) error{}

func Proxy() {
	director := func(request *http.Request) {
		request.URL.Scheme = consts.Scheme
		request.URL.Host = consts.Port
	}

	// TODO ModifyResponseがエラーの場合の挙動調べる
	modifyResponse := func(response *http.Response) error {

		if f, ok := InsertionFuncs[response.StatusCode]; ok {
			buf := new(bytes.Buffer)
			reader := io.TeeReader(response.Body, buf)
			body, err := ioutil.ReadAll(reader)
			if err != nil {
				log.Fatal(err.Error())
			}
			response.Body = ioutil.NopCloser(buf)

			f(response)

			afterBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			if len(afterBody) == 0 {
				response.Body = ioutil.NopCloser(strings.NewReader(string(body)))
			}
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
	InsertionFuncs[i] = f
}

func (insertionFuncs) Statusok(f func(response *http.Response) error) {
	addMethod(http.StatusOK, f)
}

func (insertionFuncs) StatusNotFound(f func(response *http.Response) error) {
	addMethod(http.StatusNotFound, f)
}
