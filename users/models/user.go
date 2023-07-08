package models

import "Chat/base"

type UserRole int

const (
	Admin  UserRole = 0
	Normal UserRole = 1
)

// User
//
// Model represents the User
type User struct {
	base.BaseModel
	Name     string   `json:"name" gorm:"unique"`       //Name of the user
	Email    string   `json:"email" gorm:"unique"`      //email of the user
	Role     UserRole `json:"role" gorm:"type:int"`     //role of the user
	PassWord string   `json:"-" gorm:"column:password"` //password of the user
}
