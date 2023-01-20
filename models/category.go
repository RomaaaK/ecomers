package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	Children []SubCategory
}

type SubCategory struct {
	gorm.Model
	Name       string
	CategoryID uint
}
