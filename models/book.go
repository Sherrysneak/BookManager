package models

type Book struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Quantity int
	Code     string `gorm:"uniqueIndex;size:100"`
}
