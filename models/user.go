package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `form:"email"`
	Password string
	Name     string `form:"name"`
	Role     string
	Avatar   AvatarImageStorage `media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time

	// Accepts
	AcceptPrivate bool `form:"accept-private"`
	AcceptLicense bool `form:"accept-license"`
	AcceptNews    bool `form:"accept-news"`
}

func (user User) DisplayName() string {
	return user.Email
}

func (user User) AvailableLocales() []string {
	return []string{"en-US"}
}
