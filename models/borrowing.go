package models

import "time"

type Borrowing struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	BookID     uint
	BorrowedAt time.Time
	ReturnedAt *time.Time
}
