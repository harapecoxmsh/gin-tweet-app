package services

import (
	"gin-tweet-app/models"
	"gin-tweet-app/repositories"
)

type IPostService interface {
	FindAll() (*[]models.Post, error)
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
