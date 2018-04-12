package content

import (
	"errors"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// BasicAuth adds HTTP Basic Auth check for requests that should implement it
func BasicAuth(next buffalo.Handler, backupUser, backupPass string) buffalo.Handler {
	forbiddenErr := errors.New("forbidden")
	return func(c buffalo.Context) error {
		if backupUser == "" || backupPass == "" {
			return c.Error(http.StatusForbidden, forbiddenErr)
		}

		user, password, ok := c.Request().BasicAuth()
		if !ok {
			return c.Error(http.StatusForbidden, forbiddenErr)
		}

		if backupUser != user || backupPass != password {
			return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
		}
		return next(c)
	}
}
