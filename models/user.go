package models

import (
	"github.com/jinzhu/gorm"
)

// User is a user in the system
type User struct {
	gorm.Model
	Email                string `json:"email" db:"email"`
	PasswordHash         string `json:"-" db:"password_hash"`
	Password             string `json:"-" db:"-"`
	PasswordConfirmation string `json:"-" db:"-"`
	Admin                bool   `json:"admin" db:"admin"`
}

func (u *User) DisplayName() string {

	return "Some User"
}
