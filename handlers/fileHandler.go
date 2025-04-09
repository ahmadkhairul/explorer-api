package handlers

import (
	"explorer-api/database"
	"explorer-api/models"
)

func FindFiles(fileId string, fileName string, fileType string, fileParentID string) ([]models.File, int) {
	var files []models.File
	query := database.DB
	if fileId != "" {
		query = query.Where("id = ?", fileId)
	}

	if fileParentID != "" {
		query = query.Where("parent_id = ?", fileParentID)
	}

	if fileName != "" {
		query = query.Where("name ILIKE ?", "%"+fileName+"%")
	}

	if fileType != "" {
		query = query.Where("type = ?", fileType)
	}

	query.Find(&files)

	return files, 200
}
