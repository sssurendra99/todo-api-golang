package db

import (
	"Todo-api-Golang/core/domain/entities"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGormDB() (db *gorm.DB, err error){
	DB_URL := os.Getenv("DATABASE_URL")

	db, err = gorm.Open(postgres.New(postgres.Config{
        DSN:                  DB_URL, // Data source name
        PreferSimpleProtocol: true, // Use the simple protocol
        DriverName:           "pgx", // Explicitly specify the pgx driver
    }), &gorm.Config{})

	if err != nil {
		log.Printf("failed to connect database: %v\n", err);
	}	

	db.AutoMigrate(&entities.Todo{})

	log.Println("Connecte to the database!...")

	return db, err
}