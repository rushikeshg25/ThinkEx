package database

import "time"

type User struct {
	Id        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"unique;not null"`
	Role      UserRole	`gorm:"type:user_role;not null"`
	Picture   string
	Balance   float64
	Events    []Event    `gorm:"many2many:event_participants"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


type Event struct {
	ID           string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventID      int          `gorm:"unique;autoIncrement"`
	Description  string       `gorm:"not null"`
	Title        string       `gorm:"not null"`
	Start    	 time.Time    `gorm:"not null"`
	End     	 time.Time    `gorm:"not null"`
	Status       EventStatus  `gorm:"type:event_status;default:ONGOING"`
	CreatedAt    time.Time    `gorm:"default:now()"`
	UpdatedAt    time.Time    `gorm:"default:now()"`
	Participants []User       `gorm:"many2many:event_participants"`
}


type UserRole string
const (
	ADMIN     UserRole = "admin"
	CANDIDATE UserRole = "candidate"
)

type EventStatus string
const (
	Ongoing EventStatus = "ONGOING"
	Ended   EventStatus = "ENDED"
)