package services

import (
	"errors"
	"example/ecomers/models"
)

func GetPosts() []models.Post {
	var posts []models.Post

	DB.Find(&posts)

	return posts
}

func GetPostById(id int) (models.Post, error) {
	var post models.Post

	DB.First(&post, id)

	if post.ID == 0 {
		return models.Post{}, errors.New("Post not found")
	}

	return post, nil
}
