package repositories

import (
	"gin-tweet-app/models"

	"gorm.io/gorm"
)

type IPostRepository interface {
	FindAll() (*[]models.Post, error)
	Create(newPost models.Post) (*models.Post, error)
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
func (r *PostMemoryRepository) Create(newPost models.Post) (*models.Post, error) {
	newPost.ID = uint(len(r.posts) + 1)
	r.posts = append(r.posts, newPost)
	return &newPost, nil
}

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) FindAll() (*[]models.Post, error) {
	var items []models.Post
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}
func (r *PostRepository) Create(newPost models.Post) (*models.Post, error) {
	result := r.db.Create(&newPost)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newPost, nil
}
func NewPostRepository(db *gorm.DB) IPostRepository {
	return &PostRepository{db: db}
}
