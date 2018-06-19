package admin

import (
	"github.com/markbates/goth"
)

// qorCurrentUser implements qor.CurentUser
//
// I didn't see anything in the qor suite of packages that actually implements this.
// I didn't look that deep because this is pretty easy to implement
type qorCurrentUser struct {
	displayName string
}

func (qu qorCurrentUser) DisplayName() string {
	return qu.displayName
}

func qorCurrentUserFromGothUser(usr goth.User) qorCurrentUser {
	displayName := usr.Email
	if displayName == "" {
		displayName = usr.Name
	}
	if displayName == "" {
		displayName = usr.UserID
	}
	return qorCurrentUser{displayName: displayName}
}
