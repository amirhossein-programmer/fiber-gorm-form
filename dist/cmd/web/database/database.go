package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sixfwa/fiber-api/cmd/web/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("./internal/api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database\n", err.Error())
		os.Exit(2)
	}
	log.Println("conected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	// TODO: Add migrations
	db.AutoMigrate(&models.User{})

	Database = DbInstance{Db: db}
}
