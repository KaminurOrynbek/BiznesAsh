package handler

import (
	"context"
	"net/http"

	userpb "github.com/KaminurOrynbek/BiznesAsh_lib/proto/auto-proto/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, client userpb.UserServiceClient) {
	r.POST("/register", func(c *gin.Context) {
		var req userpb.RegisterRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := client.Register(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})
}
