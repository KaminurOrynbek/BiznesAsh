package http

import (
	"ApiGateway/internal/client"
	"ApiGateway/internal/config"
	"net/http"
)

type Handler struct {
	UserClient         *client.UserClient
	ContentClient      *client.ContentClient
	NotificationClient *client.NotificationClient
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		UserClient:         &client.UserClient{BaseURL: cfg.UserServiceURL},
		ContentClient:      &client.ContentClient{BaseURL: cfg.ContentServiceURL},
		NotificationClient: &client.NotificationClient{BaseURL: cfg.NotificationServiceURL},
	}
}

func (h *Handler) UserProxy(w http.ResponseWriter, r *http.Request) {
	h.UserClient.Proxy(w, r)
}

func (h *Handler) ContentProxy(w http.ResponseWriter, r *http.Request) {
	h.ContentClient.Proxy(w, r)
}

func (h *Handler) NotificationProxy(w http.ResponseWriter, r *http.Request) {
	h.NotificationClient.Proxy(w, r)
}
