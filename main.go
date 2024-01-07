package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SantiagoBobrik/markvim-api/database"
	"github.com/SantiagoBobrik/markvim-api/routes"
	"github.com/gorilla/mux"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}

func main() {
	r := mux.NewRouter()
	err := database.InitializeDatabase()

	if err != nil {
		panic(err)
	}

	documentRoutes := routes.Routes{}
	documentRoutes.InitializeRoutes(r)

	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)

}
