package service

import "server-go/internal/database"

func GetEvent(EventID int){

}

func RegisterEvent(event database.Event){
	database.New().GetDbORM().DB().Create(&event)
}