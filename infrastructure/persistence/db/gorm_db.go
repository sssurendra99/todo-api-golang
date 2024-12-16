package db

import (
	"Todo-api-Golang/core/domain/entities"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGormDB() (db *gorm.DB, err error){
	dsn := os.Getenv("DATABASE_URL")

	dsn = "host=localhost user=postgres password=P@99!db!!73 dbname=TodoDB port=5432 sslmode=disable TimeZone=Asia/Colombo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("failed to connect database: %v\n", err);
	}	

	db.AutoMigrate(&entities.Todo{})

	log.Println("Connected to the database!...")

	return db, err
}