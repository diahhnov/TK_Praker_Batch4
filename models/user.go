package models

type Users struct {
	ID       int    `gorm:"primaryKey autoIncrement" json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
