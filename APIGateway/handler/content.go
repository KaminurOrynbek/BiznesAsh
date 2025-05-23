package handler

import (
	"context"
	contentpb "github.com/KaminurOrynbek/BiznesAsh_lib/proto/auto-proto/content"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterContentRoutes(r *gin.Engine, client contentpb.ContentServiceClient) {
	content := r.Group("/content")

	content.POST("/posts", func(c *gin.Context) {
		var req contentpb.CreatePostRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.CreatePost(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	content.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := client.GetPost(context.Background(), &contentpb.PostIdRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})
}
