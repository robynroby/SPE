package database

import (
	"fmt"
	"log"
	"os"

	"github.com/robynroby/SPE/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connect() {

	dsn := fmt.Sprintf("host=db user=%s password=%s  port=5232 TimeZone=EastAfrica/Nairobi",
		os.Getenv("DB_URL"),
		os.Getenv("DB_PASSWORD"),
		// os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect to database.\n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	}
}
