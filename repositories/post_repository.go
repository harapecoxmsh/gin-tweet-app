package repositories

import (
	"errors"
	"gin-tweet-app/models"

	"gorm.io/gorm"
)

type IPostRepository interface {
	FindAll() (*[]models.Post, error)
	FindByID(postId uint) (*models.Post, error)
	Create(newPost models.Post) (*models.Post, error)
	Delete(postId uint) error
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

func (r *PostMemoryRepository) FindByID(postId uint) (*models.Post, error) {
	for _, v := range r.posts {
		if v.ID == postId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (r *PostMemoryRepository) Create(newPost models.Post) (*models.Post, error) {
	newPost.ID = uint(len(r.posts) + 1)
	r.posts = append(r.posts, newPost)
	return &newPost, nil
}

func (r *PostMemoryRepository) Delete(postId uint) error {
	for i, v := range r.posts {
		if v.ID == postId {
			r.posts = append(r.posts[:i], r.posts[i+1:]...)
		}
	}
	return errors.New("Item not found")
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

func (r *PostRepository) FindByID(postId uint) (*models.Post, error) {
	var post models.Post
	result := r.db.First(&post, "id = ?", postId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}
		return nil, result.Error
	}
	return &post, nil
}

func (r *PostRepository) Create(newPost models.Post) (*models.Post, error) {
	result := r.db.Create(&newPost)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newPost, nil
}

func (r *PostRepository) Delete(postId uint) error {
	deletePost, err := r.FindByID(postId)
	if err != nil {
		return err
	}
	result := r.db.Delete(&deletePost)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func NewPostRepository(db *gorm.DB) IPostRepository {
	return &PostRepository{db: db}
}
