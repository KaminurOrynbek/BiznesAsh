package client

import (
	"io"
	"net/http"
)

type NotificationClient struct {
	BaseURL string
}

func (c *NotificationClient) Proxy(w http.ResponseWriter, r *http.Request) {
	proxyReq, _ := http.NewRequest(r.Method, c.BaseURL+r.URL.Path, r.Body)
	proxyReq.Header = r.Header
	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, "Notification service unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
