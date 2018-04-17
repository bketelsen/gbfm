package content

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/nilslice/jwt"
)

type loginHandlers struct {
	successRedir string // where to redirect when a login succeeds
	showFormPath string // the path that renders the login form
}

// GET /admin/login
func (l loginHandlers) showForm(c buffalo.Context) error {
	if IsValid(c) {
		return c.Redirect(http.StatusFound, l.successRedir)
	}
	u := new(models.User)
	c.Set("user", u)
	return c.Render(http.StatusOK, r.HTML("auth/new.html"))
}

// POST /admin/login
func (l loginHandlers) try(c buffalo.Context) error {
	if IsValid(c) {
		return c.Redirect(http.StatusFound, l.successRedir)
	}

	str := &struct {
		Email    string
		Password string
	}{}
	if err := c.Bind(str); err != nil {
		fmt.Println("error binding post")
		c.Flash().Add("error", "email or password missing")
		return c.Redirect(http.StatusUnauthorized, l.showFormPath)
	}
	// check email & password
	usr := new(models.User)
	if err := models.DB.Where("email = ?", str.Email).First(usr); err != nil {
		fmt.Println("error with sql statement", err)
		c.Flash().Add("error", "email not found")
		return c.Redirect(http.StatusFound, l.showFormPath)
	}
	if usr == nil {
		fmt.Println("empty user")
		c.Flash().Add("error", "no such user")
		return c.Redirect(http.StatusFound, l.showFormPath)
	}

	if !IsUser(usr, str.Password) {
		fmt.Println("password check fails")
		c.Flash().Add("error", "invalid password")
		return c.Redirect(http.StatusFound, l.showFormPath)
	}
	// create new token
	week := time.Now().Add(time.Hour * 24 * 7)
	claims := map[string]interface{}{
		"exp":  week,
		"user": usr.Email,
	}
	token, err := jwt.New(claims)
	if err != nil {
		log.Println(err)
		c.Flash().Add("error", "couldn't create token")
		return c.Redirect(http.StatusFound, l.showFormPath)
	}
	c.Session().Set("_token", token)
	c.Session().Save()

	return c.Redirect(http.StatusFound, l.successRedir)
}
