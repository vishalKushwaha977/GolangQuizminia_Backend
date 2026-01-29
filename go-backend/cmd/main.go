package main

import (
	"log"
	"go-backend/internal/db"
	"go-backend/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {
	db.ConnectDB()
	r := gin.Default()
	// CORS
	r.Use(cors.Default())
	// Routes
	routes.SetupRoutes(r)
	log.Println("Server started on :8080")
	r.Run(":8080")
}
