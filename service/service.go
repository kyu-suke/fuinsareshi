package service

import (
	"log"
	"net/http"
	"net/http/httputil"
	//"fmt"
	//"io/ioutil"
)

func Proxy() {

	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":8080"
	}
	modifyResponse := func(response *http.Response) error {
		//r, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(r))
		//fmt.Println(response.StatusCode)
		//for k, v := range response.Header {
		//	fmt.Println(k, v)
		//}
		return nil
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
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
