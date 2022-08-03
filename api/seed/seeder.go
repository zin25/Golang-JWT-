package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"project/api/models"
)


//DUMMY DATA
var users = []models.User{
	models.User{
		Fullname: "Steven Kurniawan",
		Email:    "Kurniawan@gmail.com",
		Phone: "087934986918",
		Password: "password",
	},
	models.User{
		Fullname: "Martin Garut",
		Email:    "Garut@gmail.com",
		Phone: "78910491704",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Desc: "Hello world 1",
		Types: "Artikel",
	},
	models.Post{
		Title:   "Title 2",
		Desc: "Hello world 2",
		Types: "Idea",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}