package services

import (
	"example/ecomers/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")

	dsn := db_user + ":" + db_pass + "@tcp(localhost:3306)/ecomers?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&models.Brand{}, &models.Category{}, &models.Product{}, &models.Blog{})

	var brands []models.Brand
	var category []models.Category
	var products []models.Product
	var blogs []models.Blog

	db.Find(&brands)
	db.Model(&models.Category{}).Preload("Childrens").Find(&category)
	db.Find(&products)
	db.Find(&blogs)

	if len(blogs) == 0 {
		post1 := models.Blog{
			Title: "GIRLS PINK T SHIRT ARRIVED IN STORE",
			Image: "images/blog/blog-one.jpg",
			Text: `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

			Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.
			
			Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.
			
			Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.
			`,
		}

		post2 := models.Blog{
			Title: "GIRLS PINK T SHIRT ARRIVED IN STORE",
			Image: "images/blog/blog-two.jpg",
			Text: `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

			Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.
			
			Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.
			
			Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.
			`,
		}

		post3 := models.Blog{
			Title: "GIRLS PINK T SHIRT ARRIVED IN STORE",
			Image: "images/blog/blog-three.jpg",
			Text: `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

			Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.
			
			Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt.
			
			Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.
			`,
		}

		db.Create(&post1)
		db.Create(&post2)
		db.Create(&post3)

	}

	if len(category) == 0 {

		nike := models.Category{Name: "Nike"}
		adidas := models.Category{Name: "Adidas"}
		zara := models.Category{Name: "Zara"}

		sportsware := models.Category{Name: "SPORTSWEAR", Childrens: []models.Category{nike, adidas, zara}}
		mens := models.Category{Name: "Mens", Childrens: []models.Category{zara, adidas}}
		womens := models.Category{Name: "Womens", Childrens: []models.Category{adidas, nike}}
		other := models.Category{Name: "Other"}

		db.Create(&sportsware)
		db.Create(&mens)
		db.Create(&womens)
		db.Create(&other)
	}

	if len(brands) == 0 {
		nike := models.Brand{Name: "Nike"}
		zara := models.Brand{Name: "Zara"}
		adidas := models.Brand{Name: "Adidas"}
		uniqlo := models.Brand{Name: "Uniqlo"}

		db.Create(&nike)
		db.Create(&zara)
		db.Create(&adidas)
		db.Create(&uniqlo)
	}

	if len(products) == 0 {
		prod1 := models.Product{
			Name:       "Anne Klein Sleeveless Colorblock Scuba",
			Rating:     4,
			Price:      59,
			Quantity:   3,
			Condition:  "New",
			Image:      "/images/home/product1.jpg",
			BrandID:    &brands[0].ID,
			CategoryID: &category[10].ID,
		}
		prod2 := models.Product{
			Name:       "Anne Klein Sleeveless Colorblock Scuba",
			Rating:     4,
			Price:      59,
			Quantity:   3,
			Condition:  "New",
			Image:      "/images/home/product1.jpg",
			BrandID:    &brands[0].ID,
			CategoryID: &category[10].ID,
		}
		prod3 := models.Product{
			Name:       "Anne Klein Sleeveless Colorblock Scuba",
			Rating:     4,
			Price:      59,
			Quantity:   3,
			Condition:  "New",
			Image:      "/images/home/product1.jpg",
			BrandID:    &brands[0].ID,
			CategoryID: &category[10].ID,
		}
		prod4 := models.Product{
			Name:       "Anne Klein Sleeveless Colorblock Scuba",
			Rating:     4,
			Price:      59,
			Quantity:   3,
			Condition:  "New",
			Image:      "/images/home/product1.jpg",
			BrandID:    &brands[1].ID,
			CategoryID: &category[10].ID,
		}
		prod5 := models.Product{
			Name:       "Anne Klein Sleeveless Colorblock Scuba",
			Rating:     4,
			Price:      59,
			Quantity:   3,
			Condition:  "New",
			Image:      "/images/home/product1.jpg",
			BrandID:    &brands[1].ID,
			CategoryID: &category[10].ID,
		}
		prod6 := models.Product{
			Name:       "Anne Klein Sleeveless Colorblock Scuba",
			Rating:     4,
			Price:      59,
			Quantity:   3,
			Condition:  "New",
			Image:      "/images/home/product1.jpg",
			BrandID:    &brands[1].ID,
			CategoryID: &category[10].ID,
		}

		db.Create(&prod1)
		db.Create(&prod2)
		db.Create(&prod3)
		db.Create(&prod4)
		db.Create(&prod5)
		db.Create(&prod6)
	}

	DB = db
}
