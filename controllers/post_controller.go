package controllers

import (
	"gin-tweet-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	FindAll(ctx *gin.Context)
}

type PostController struct {
	service services.IPostService
}

func NewPostController(service services.IPostService) IPostController {
	return &PostController{service: service}
}

func (c *PostController) FindAll(ctx *gin.Context) {
	posts, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}
