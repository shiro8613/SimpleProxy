package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func HttpProxy(addr string, port int, balanceHost []string, Proto string) {
	director := func(request *http.Request) {
		request.URL.Scheme = Proto
		request.URL.Host = ""

	}

	rp := &httputil.ReverseProxy{Director: director}

	server := http.Server{
		Addr: fmt.Sprintf("%s:%d", addr, port),
		Handler: rp,
	}
}