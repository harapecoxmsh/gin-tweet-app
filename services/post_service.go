package services

import (
	"gin-tweet-app/dto"
	"gin-tweet-app/models"
	"gin-tweet-app/repositories"
)

type IPostService interface {
	FindAll() (*[]models.Post, error)
	FindByID(postId uint) (*models.Post, error)
	Create(postInput dto.PostInput) (*models.Post, error)
	Delete(postId uint) error
}

type PostService struct {
	repository repositories.IPostRepository
}

func NewPostService(repository repositories.IPostRepository) IPostService {
	return &PostService{repository: repository}
}

func (s *PostService) FindAll() (*[]models.Post, error) {
	return s.repository.FindAll()
}

func (s *PostService) FindByID(postId uint) (*models.Post, error) {
	return s.repository.FindByID(postId)
}

func (s *PostService) Delete(postId uint) error {
	return s.repository.Delete(postId)
}

func (s *PostService) Create(postInput dto.PostInput) (*models.Post, error) {
	newPost := models.Post{
		Content: postInput.Content,
	}
	return s.repository.Create(newPost)
}
