package database

import "time"

type User struct {
	Id        string `gorm:"primaryKey"` 
	Role      string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

