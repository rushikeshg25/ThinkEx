package database

import "time"

type User struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"unique;not null"`
	Role      UserRole
	Picture   string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRole string

const (
	ADMIN     UserRole = "admin"
	CANDIDATE UserRole = "candidate"
)
