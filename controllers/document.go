package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SantiagoBobrik/markvim-api/models"
	"github.com/gorilla/mux"
)

type DocumentBody struct {
	Base64 string `json:"base64"`
}

type Document struct{}

func (d *Document) GetDocument(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	document, err := models.GetDocument(id)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(document)
}

func (d *Document) GetDocuments(w http.ResponseWriter, r *http.Request) {

	documents, err := models.GetAllDocuments()

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(documents)
}

func (d *Document) AddDocument(w http.ResponseWriter, r *http.Request) {

	body := &DocumentBody{}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		panic(err)
	}

	if body.Base64 == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = models.AddDocument(body.Base64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)

}
