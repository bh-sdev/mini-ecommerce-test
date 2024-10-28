// main.go
package main

import (
	"log"
	"net/http"

	"github.com/superdev/ecommerce/gateway/config"
	"github.com/superdev/ecommerce/gateway/database"
	"github.com/superdev/ecommerce/gateway/handlers"
	"github.com/superdev/ecommerce/gateway/middleware"

	"github.com/gorilla/mux"
)

func main() {

	config.InitConfig()
	db, err := database.NewDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	r := mux.NewRouter()
	r.Use(middleware.JwtAuthMiddleware)

	r.HandleFunc("/products", handlers.ListProducts).Methods("GET")
	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")

	log.Println("API Gateway running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
