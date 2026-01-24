package main

import (
	"log"
	"net/http"
	"go-backend/internal/db"
	"go-backend/internal/routes"
)

func main() {
	db.ConnectDB()
	r := routes.SetupRoutes()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
