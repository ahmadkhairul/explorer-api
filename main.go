package main

import (
	"explorer-api/database"
	"explorer-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Koneksi ke PostgreSQL
	database.ConnectDB()

	// Setup Fiber
	app := fiber.New()

	// Setup Routes
	routes.SetupRoutes(app)

	// Jalankan server di port 3000
	app.Listen(":3000")
}
