package model

import (
	"gorm.io/gorm"
	// "time"
)

type User struct {
	gorm.Model
	Username	string		`json:"username" gorm:"type:varchar(255);not null"`
	Fullname	string		`json:"fullname" gorm:"type:varchar(255); not null"`
	Age			int			`json:"age" gorm:"type:int"`
	Password	string		`json:"password" gorm:"type:varchar(255); not null"`
		
}

type UserResponse struct{
	Name		string
	Message		string
}
