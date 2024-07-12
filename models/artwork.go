package models

type Artwork struct {
	ID          uint   `gorm:"primaryKey"`
	MuseumID    uint   `json:"museum_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       []byte `json:"image"`
	Author      string `json:"author"`
	Active      bool   `json:"active"`
}
