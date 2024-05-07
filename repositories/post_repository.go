package repositories

import "gin-tweet-app/models"

type IPostRepository interface {
	FindAll() (*[]models.Post, error)
}

type PostMemoryRepository struct {
	posts []models.Post
}

func NewPostMemoryRepository(posts []models.Post) IPostRepository {
	return &PostMemoryRepository{posts: posts}
}

func (r *PostMemoryRepository) FindAll() (*[]models.Post, error) {
	return &r.posts, nil
}
