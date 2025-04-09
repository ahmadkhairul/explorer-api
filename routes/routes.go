package routes

import (
	"explorer-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/files", controllers.GetFile)
	api.Get("/folder/:parent_id", controllers.GetFileByParentID)
	api.Get("/files/:id", controllers.GetFileByID)
	api.Post("/files", controllers.CreateFile)
}
