package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func Connect() {
	connectionString := os.Getenv("DB_URL")

	conn, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	err = conn.Ping()

	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	Db = conn

	println("DB Connection successful")
}
