package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gophersnacks/gbfm/models"
	"github.com/pkg/errors"
)

// UsersNew shows the new user form
func UsersNew(c buffalo.Context) error {
	u := models.User{}
	c.Set("user", u)
	return c.Render(200, r.HTML("users/new.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	// tx := c.Value("tx").(*pop.Connection)
	if err := u.Create(models.GORM); err != nil {
		return errors.WithStack(err)
	}

	// if verrs.HasAny() {
	// 	c.Set("user", u)
	// 	c.Set("errors", verrs)
	// 	return c.Render(200, r.HTML("users/new.html"))
	// }

	c.Session().Set("current_user_id", u.ID)
	return c.Render(200, r.JSON(u))

}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
			c.Set("current_user_id", u.ID)
			fmt.Println("Current User", u)
			if u.Admin {
				c.Session().Set("admin", "true")
				c.Set("admin", true)
			}
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Error(403, errors.New("Forbidden"))
		}
		return next(c)
	}
}

// AdminAuthorize require a user be logged in before accessing a route
func AdminAuthorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")

			return c.Error(403, errors.New("Forbidden"))
		}
		if admin := c.Session().Get("admin"); admin == nil {
			c.Flash().Add("danger", "Nice try, but you're not the boss of me.")

			return c.Error(403, errors.New("Forbidden"))
		}
		return next(c)
	}
}
