package client

import (
	"io"
	"net/http"
)

type UserClient struct {
	BaseURL string
}

func (c *UserClient) Proxy(w http.ResponseWriter, r *http.Request) {
	proxyReq, _ := http.NewRequest(r.Method, c.BaseURL+r.URL.Path, r.Body)
	proxyReq.Header = r.Header
	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, "User service unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
