package inits

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBinit() {
	dbUrl := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	DB = db

}
