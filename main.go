package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"example.com/m/v2/database"
	"example.com/m/v2/database/migrations"
	"example.com/m/v2/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	database.InitDatabase()
	migrations.RunMigrations()

	app := fiber.New()

	route.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
