package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// -- open items
	// need to get these creds from env 

	dsn := "root:Suresh02!@tcp(localhost:3306)/sys"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting the database..!")
	}

	DB = db

	log.Print("connection to database is sucessfull..!")

}
