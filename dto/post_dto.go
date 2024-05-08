package dto

type PostInput struct {
	Content string `json:"content" binding:"required,gt=1,lte=128"`
}
