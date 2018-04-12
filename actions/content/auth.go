// Package user contains the basic admin user creation and authentication code,
// specific to Ponzu systems.
package content

import (
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/nilslice/jwt"
)

// Auth is HTTP middleware to ensure the request has proper token credentials
func Auth(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if IsValid(c) {
			return next(c)
		}
		return c.Redirect(http.StatusFound, "/admin/login")
	}
}

// IsValid checks if the user request is authenticated
func IsValid(c buffalo.Context) bool {
	tknIface := c.Session().Get("_token")
	if tknIface == nil {
		return false
	}
	tkn, ok := tknIface.(string)
	if !ok {
		return false
	}
	return jwt.Passes(tkn)
}

// IsUser checks for consistency in email/pass combination
func IsUser(usr *models.User, password string) bool {
	if err := usr.ComparePassword(password); err != nil {
		log.Println("Error checking password:", err)
		return false
	}

	return true
}
