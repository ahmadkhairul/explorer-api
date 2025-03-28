package models

type File struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Size     int    `json:"size" gorm:"default: 0" `
	ParentID string `json:"parent_id" gorm:"default: null" `
	Path     string `json:"path" gorm:"default: null" `
}
