package services

import (
	"errors"
	"example/ecomers/models"
)

func GetBlogs() []models.Blog {
	var posts []models.Blog

	DB.Find(&posts)

	return posts
}

func GetBlogById(id int) (models.Blog, error) {
	var post models.Blog

	DB.First(&post, id)

	if post.ID == 0 {
		return models.Blog{}, errors.New("blog not found")
	}

	return post, nil
}
