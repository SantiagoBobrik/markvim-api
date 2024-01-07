package models

import (
	"time"

	"github.com/SantiagoBobrik/markvim-api/database"
	"github.com/google/uuid"
)

// TODO -  Implement the Document struct
type Document struct {
	ID        string    `json:"id"`
	Base64    string    `json:"base64"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DocumentResponse struct {
	ID     string `json:"id"`
	Base64 string `json:"base64"`
}

func GetDocument(id string) (*DocumentResponse, error) {

	document := &DocumentResponse{}
	err := database.DB.QueryRow("SELECT id, base64 FROM documents WHERE id = ?", id).Scan(&document.ID, &document.Base64)

	if err != nil {
		return nil, err
	}

	return document, nil
}

func AddDocument(base64 string) error {

	id := uuid.New().String()

	statement, err := database.DB.Prepare(
		"INSERT INTO documents (id, base64, created_at, updated_at) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(
		id,
		base64,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func GetAllDocuments() ([]DocumentResponse, error) {

	rows, err := database.DB.Query("SELECT id, base64 FROM documents")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	documents := []DocumentResponse{}

	for rows.Next() {
		var document DocumentResponse
		err := rows.Scan(&document.ID, &document.Base64)

		if err != nil {
			return nil, err
		}

		documents = append(documents, document)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}
