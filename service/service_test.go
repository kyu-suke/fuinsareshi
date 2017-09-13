package service

import (
	"io/ioutil"
	"testing"

	"fmt"

	"net/http"
)

func TestProxy(t *testing.T) {

	go Proxy()

	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	go http.ListenAndServe(":8080", nil)

	c := http.Client{}
	resp, err := c.Get("http://localhost:8080")
	if err != nil {
		t.Fatal("url変です")
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("レスポンス変です")
	}

	proxyClient := http.Client{}
	proxyResp, err := proxyClient.Get("http://localhost:9000")
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
