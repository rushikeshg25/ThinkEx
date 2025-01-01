package service

import (
	"server-go/internal/database"
	"time"
)

func Getuser(Email string, AvatarURL string){
	database.New().GetDbORM().Create(&database.User{
		Email:     Email,
		Role:      database.CANDIDATE,
		Picture:   AvatarURL,
		Balance:   100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func RegisterUser(Email string, AvatarURL string){
	database.New().GetDbORM().Create(&database.User{
		Email:     Email,
		Role:      database.CANDIDATE,
		Picture:   AvatarURL,
		Balance:   100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}