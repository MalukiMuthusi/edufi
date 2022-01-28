package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var Db *sql.DB

func StartDB() {

	dbUser := os.Getenv("user")
	dbPwd := os.Getenv("pwd")
	dbName := os.Getenv("name")
	dbPort := os.Getenv("port")
	dbHost := os.Getenv("host")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbHost, dbUser, dbPwd, dbPort, dbName)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Db = db

	// Confirm a successful connection.
	if err := Db.Ping(); err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}

}
