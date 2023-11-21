package database

import (
	"fmt"
	"github.com/Zulhaditya/restful-note/config"
	"github.com/Zulhaditya/restful-note/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// declare variable for database
var DB *gorm.DB

// connect to database
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		log.Println("Error")
	}

	// connection url to connect postgresql database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	// initialize DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Connection open to database...")
}
