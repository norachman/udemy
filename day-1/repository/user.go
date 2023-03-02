package repository

import (
	"day-1/db"
	"day-1/model"
)

func GetAllUser() []model.User {
	var user []model.User 

	data := db.GetConnectionDB()
	data.Find(&user)
	
	return user
}

func GetUserById(id int) model.User {
	var user model.User

	data := db.GetConnectionDB()
	data.Find(&user)

	return user
}
