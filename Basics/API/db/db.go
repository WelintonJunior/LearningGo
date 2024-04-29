package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// INICIAR O BANCO
func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:2206@/escola")

	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	createTables()
}

// CRIAR AS TABELAS (DESNECESSARIO)
func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INT
	)
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(err)
	}
}
