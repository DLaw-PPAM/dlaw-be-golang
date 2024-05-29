package migrations

import (
	"fmt"

	"example.com/m/v2/database"
	"example.com/m/v2/model/entity"

	"log"
)

func RunMigrations() {
	if database.DB == nil {
		fmt.Printf("Database connection: %v\n", database.DB)
		log.Fatal("Database connection is nil")
	}

	err := database.DB.AutoMigrate(&entity.User{}, &entity.LawyerUser{}, &entity.Specialties{}, &entity.Case{}, &entity.Review{}, &entity.DetailLawyer{})

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	fmt.Println("Migration run successfully")
}
