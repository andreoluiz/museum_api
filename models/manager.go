package models

type Manager struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Active    bool   `json:"active"`
}

type ManagerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
