package routes

import (
	"explorer-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/files", controllers.GetFile)
	api.Get("/files/:parent_id", controllers.GetFileByID)
	api.Post("/files", controllers.CreateFile)
}
