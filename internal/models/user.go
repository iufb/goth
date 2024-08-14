package models

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
}
