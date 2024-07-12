package models

type Museum struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       []byte    `json:"image" gorm:"type:bytea"`
	Category1   string    `json:"category1"`
	Category2   string    `json:"category2"`
	Link        string    `json:"link"`
	Address     string    `json:"address"`
	CEP         string    `json:"cep"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Information string    `json:"information"`
	ManagerID   uint      `json:"manager_id"`
	Artworks    []Artwork `json:"artworks"`
	Active      bool      `json:"active"`
}
