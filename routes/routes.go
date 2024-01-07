package routes

import (
	"net/http"

	"github.com/SantiagoBobrik/markvim-api/controllers"
	"github.com/gorilla/mux"
)

type Routes struct{}

func (routes *Routes) InitializeRoutes(r *mux.Router) {
	/* ----------------------------- Document routes ---------------------------- */
	documentController := controllers.Document{}
	r.HandleFunc("/document/{id}", documentController.GetDocument).Methods(http.MethodGet)
	r.HandleFunc("/documents", documentController.GetDocuments).Methods(http.MethodGet)
	r.HandleFunc("/document", documentController.AddDocument).Methods(http.MethodPost)
	/* -------------------------------------------------------------------------- */
}
