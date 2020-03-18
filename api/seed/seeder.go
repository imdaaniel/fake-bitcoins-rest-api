package seed

import (
	"log"

	"github.com/imdaaniel/bitcoins-rest-api/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Name:        "Adam Levine",
		Email:       "adam.levine@gmail.com",
		Password:    "password",
		DateOfBirth: "1979-03-18",
	},
	models.User{
		Name:        "Elon Musk",
		Email:       "elonmusk@gmail.com",
		Password:    "password",
		DateOfBirth: "1971-06-28",
	},
}

var orders = []models.Order{
	models.Order{
		AuthorID: 1,
		Amount:   1.23,
		Action:   "buy",
	},
	models.Order{
		AuthorID: 2,
		Amount:   5.43,
		Action:   "buy",
	},
	models.Order{
		AuthorID: 1,
		Amount:   1.15,
		Action:   "sell",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Order{}, models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Order{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Order{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("Foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}
	}

	for i := range orders {
		err = db.Debug().Model(&models.Order{}).Create(&orders[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed orders table: %v", err)
		}
	}
}
