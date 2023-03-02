package model

import (
	"gorm.io/gorm"
	// "time"
)

type User struct {
	gorm.Model
	Name 		string		`gorm:"varchar(255);not null"`
	Age			int			`gorm:"int"`
		
}

type UserResponse struct{
	Name		string
	Message		string
}
