package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type DocumentResponse struct {
	ID     string `json:"id"`
	Base64 string `json:"base64"`
}

func InitializeDatabase() error {

	fmt.Println("Initializing database")

	var err error
	d, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return err
	}

	DB = d

	err = initializeTables()

	if err != nil {
		return err
	}

	fmt.Println("Database is running")

	return nil

}

func initializeTables() error {
	if err := createDocumentTable(); err != nil {
		return err
	}
	return nil
}

func createDocumentTable() error {

	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS documents (id varchar(255) PRIMARY KEY, base64 TEXT, created_at DATETIME, updated_at DATETIME)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(); err != nil {
		return err
	}

	fmt.Println("Tables are being created")
	return nil
}
