package controllers

import (
	"gin-tweet-app/dto"
	"gin-tweet-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
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
func (c *PostController) Create(ctx *gin.Context) {
	var input dto.PostInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newPost, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newPost})
}
