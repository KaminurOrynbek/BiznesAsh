package client

import (
	"io"
	"net/http"
)

type ContentClient struct {
	BaseURL string
}

func (c *ContentClient) Proxy(w http.ResponseWriter, r *http.Request) {
	proxyReq, _ := http.NewRequest(r.Method, c.BaseURL+r.URL.Path, r.Body)
	proxyReq.Header = r.Header
	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, "Content service unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
