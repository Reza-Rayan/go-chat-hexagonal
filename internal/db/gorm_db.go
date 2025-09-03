package db

import (
	"fmt"
	"log"

	"github.com/Reza-Rayan/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.LoadConfig()

	var err error
	DB, err = gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to db: %v", err))
	}

	// AutoMigrate tables
	err = DB.AutoMigrate(
	//	Add Models HERE
	)
	if err != nil {
		panic(fmt.Sprintf("Migration failed: %v", err))
	}

	log.Printf("✅ Connected to SQLite database: %s", cfg.Database.Path)
}
