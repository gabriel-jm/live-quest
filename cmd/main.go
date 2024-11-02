package main

import (
	"live-quest/internal/database"
	"live-quest/internal/users"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnvErr := godotenv.Load(".env")

	if loadEnvErr != nil {
		log.Fatalln("Error loading env file:", loadEnvErr)
	}

	database.Connect()
	defer database.Db.Close()

	runMigrations()

	e := echo.New()
	users.Routes(e)

	log.Println("Server on http://localhost:8000")
	e.Logger.Fatal(e.Start(":8000"))
}

func runMigrations() {
	driver, err := postgres.WithInstance(database.Db.DB, &postgres.Config{})

	if err != nil {
		log.Fatal("Error getting postgres driver instance", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal("Error initializing migrator", err)
	}

	migrator.Up()
}
