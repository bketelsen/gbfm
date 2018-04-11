package content

import (
	"errors"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/pkg/system/db"
)

// BasicAuth adds HTTP Basic Auth check for requests that should implement it
func BasicAuth(next buffalo.Handler) buffalo.Handler {
	forbiddenErr := errors.New("forbidden")
	return func(c buffalo.Context) error {
		u := db.ConfigCache("backup_basic_auth_user").(string)
		p := db.ConfigCache("backup_basic_auth_password").(string)

		if u == "" || p == "" {
			return c.Error(http.StatusForbidden, forbiddenErr)
		}

		user, password, ok := c.Request().BasicAuth()

		if !ok {
			return c.Error(http.StatusForbidden, forbiddenErr)
		}

		if u != user || p != password {
			return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
		}
		return next(c)
	}
}
