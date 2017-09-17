package service

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"fmt"

	"net/http"

	"github.com/kyu-suke/fuinsareshi/consts"
)

func TestProxy(t *testing.T) {

	go Proxy()

	http.HandleFunc("/index", handler)
	go http.ListenAndServe(consts.Port, nil)

	c := http.Client{}
	resp, err := c.Get("http://localhost:8080/index")
	if err != nil {
		t.Fatal("url変です")
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("レスポンス変です")
	}

	proxyClient := http.Client{}
	proxyResp, err := proxyClient.Get("http://localhost:9000/index")
	if err != nil {
		t.Fatal("url変です")
	}
	proxyResult, err := ioutil.ReadAll(proxyResp.Body)
	if err != nil {
		t.Fatal("レスポンス変です")
	}

	if string(r) != string(proxyResult) {
		t.Fatal("レスポンスがちがいます")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func Test中間処理加えるテスト(t *testing.T) {
	funcBody := ""
	InsertionFunc.Statusok(func(r *http.Response) error {
		buf := new(bytes.Buffer)
		reader := io.TeeReader(r.Body, buf)
		r.Body = ioutil.NopCloser(buf)
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			t.Fatal("エラーです")
		}
		funcBody = string(body)
		return nil
	})

	proxyClient := http.Client{}
	proxyResp, err := proxyClient.Get("http://localhost:9000/index")
	if err != nil {
		t.Fatal("url変です")
	}
	proxyResult, err := ioutil.ReadAll(proxyResp.Body)
	if err != nil {
		t.Fatal("レスポンス変です")
	}

	if funcBody != string(proxyResult) {
		t.Fatal("結果がおかしいです")
	}

	InsertionFunc.StatusNotFound(func(r *http.Response) error {
		buf := new(bytes.Buffer)
		reader := io.TeeReader(r.Body, buf)
		r.Body = ioutil.NopCloser(buf)
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			t.Fatal("エラーです")
		}
		funcBody = string(body)
		return nil
	})

	notFoundProxyClient := http.Client{}
	notFoundProxyResp, err := notFoundProxyClient.Get("http://localhost:9000/notfoud")
	if err != nil {
		t.Fatal("url変です")
	}
	notFoundProxyResult, err := ioutil.ReadAll(notFoundProxyResp.Body)
	if err != nil {
		t.Fatal("レスポンス変です")
	}

	if funcBody != string(notFoundProxyResult) {
		t.Fatal("結果がおかしいです")
	}
}
