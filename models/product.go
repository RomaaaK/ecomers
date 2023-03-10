package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Rating     uint
	Price      float64
	Quantity   uint
	Condition  string
	Image      string
	Brand      Brand
	BrandID    *uint
	CategoryID *uint
}
