package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	userpq := os.Getenv("userpq")
	passwordpq := os.Getenv("passwordpq")
	dbnamepq := os.Getenv("dbnamepq")
	hostpq := os.Getenv("hostpq")
	portpq := os.Getenv("portpq")
	sslmodepq := os.Getenv("sslmodepq")

	conn := fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v", userpq, passwordpq, dbnamepq, hostpq, portpq, sslmodepq)

	var err error
	DB, err = sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
