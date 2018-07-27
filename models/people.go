package models

import (
	"log"
	db "gin-demo/database"
)

type People struct {
	Id int64 `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName string `json:"last_name" form:"last_name"`
}
// 查询ID
func GetPeopleById(Id int ) (a *People){
	var people People
	has, err := db.Orm.Id(Id).Get(&people)
	if(err != nil){
		log.Fatal(err.Error())
		return nil
	}
	if has == false{
		return nil
	}
	return &people
}