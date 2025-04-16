package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	for i := 0; i < 10; i++ { // try for ~10 seconds
		DB, err = sql.Open("postgres", connStr)
		if err == nil {
			err = DB.Ping()
		}

		if err == nil {
			fmt.Println("✅ Connected to Postgres")
			return
		}

		fmt.Println("⏳ Waiting for Postgres to be ready...")
		time.Sleep(2 * time.Second)
	}

	panic("❌ Cannot connect to Postgres: " + err.Error())
}
