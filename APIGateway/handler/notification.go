package handler

import (
	"context"
	"net/http"

	notificationpb "github.com/KaminurOrynbek/BiznesAsh_lib/proto/auto-proto/notification"
	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(r *gin.Engine, client notificationpb.NotificationServiceClient) {
	notify := r.Group("/notify")

	notify.POST("/welcome", func(c *gin.Context) {
		var req notificationpb.EmailRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.SendWelcomeEmail(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	notify.POST("/system-message", func(c *gin.Context) {
		var req notificationpb.SystemMessageRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.NotifySystemMessage(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})
}
