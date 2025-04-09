package controllers

import (
	"explorer-api/database"
	"explorer-api/handlers"
	"explorer-api/models"
	"explorer-api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetFile(c *fiber.Ctx) error {
	fileName := c.Query("name")
	fileType := c.Query("type")

	files, statusCode := handlers.FindFiles("", fileName, fileType, "")

	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "OK",
		"message": "Success",
		"data":    files,
	})
}

func GetFileByParentID(c *fiber.Ctx) error {
	fileParentID := c.Params("parent_id")
	fileName := c.Query("name")
	fileType := c.Query("type")

	files, statusCode := handlers.FindFiles("", fileName, fileType, fileParentID)

	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "OK",
		"message": "Success",
		"data":    files,
	})
}

func GetFileByID(c *fiber.Ctx) error {
	fileID := c.Params("id")

	files, statusCode := handlers.FindFiles(fileID, "", "", "")

	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "OK",
		"message": "Success",
		"data":    files,
	})
}

func CreateFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	fileName := c.FormValue("name")

	if file == nil && fileName == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "File or folder name is required",
		})
	}

	newFile := models.File{
		Name: fileName,
	}

	fileParentID := c.FormValue("parent_id")

	if fileParentID != "" {
		newFile.ParentID = fileParentID
	}

	newFile.Type = "folder"
	if file != nil {
		newFile.Type = "file"
	}

	if file != nil {
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status":  "ERROR",
				"message": "Failed to upload file",
				"error":   err.Error(),
			})
		}

		newFile.Size = int(file.Size)
		newFile.Name = file.Filename
		savePath := "./uploads/" + utils.RandStr(10) + file.Filename

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status":  "ERROR",
				"message": "Failed to save file",
				"error":   err.Error(),
			})
		}
		newFile.Path = savePath
	}

	if err := database.DB.Create(&newFile).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "Failed to save file to database",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "OK",
		"message": "Success",
		"data":    newFile,
	})
}

func UpdateFile(c *fiber.Ctx) error {
	fileID := c.Params("id")
	fileName := c.FormValue("newFileName")

	if fileName == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "New File or folder name is required",
		})
	}

	updatedFile := models.File{
		Name: fileName,
	}

	if err := database.DB.Update(fileID, &updatedFile).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "Failed to update file name",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "OK",
		"message": "Success",
		"data":    updatedFile,
	})
}
