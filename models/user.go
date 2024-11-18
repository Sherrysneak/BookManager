package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:255;uniqueIndex"`
}
