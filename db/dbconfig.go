package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	
    _ "github.com/lib/pq"
)

func ConnectDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Faled to connect to db", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	fmt.Println("Database connection successful")
	return db
}