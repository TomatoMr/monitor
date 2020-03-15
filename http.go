package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewMultipleHostsReverseProxy(target url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		go func() {
			ch <- req.RequestURI
		}()
	}
	return &httputil.ReverseProxy{Director: director}
}

func Proxy() {
	InitData()
	target := url.URL{
		Scheme: "http",
		Host:   ":9091",
	}
	proxy := NewMultipleHostsReverseProxy(target)
	log.Fatal(http.ListenAndServe(":9090", proxy))

}

var ch chan string
var labels map[string]float64

func InitData() {
	ch = make(chan string) //api通道，每访问一个api，就将api路径放至此通道
	labels = make(map[string]float64) //api访问次数存储的地方
}
