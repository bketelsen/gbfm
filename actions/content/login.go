package content

import (
	"log"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/models"
	"github.com/nilslice/jwt"
)

func loginHandler(c buffalo.Context) error {
	if IsValid(c) {
		return c.Redirect(http.StatusFound, "/admin")
	}
	u := new(models.User)
	c.Set("user", u)
	return c.Render(http.StatusOK, r.HTML("auth/new.html"))
}

func attemptLoginHandler(c buffalo.Context) error {
	if IsValid(c) {
		return c.Redirect(http.StatusFound, "/admin")
	}

	str := &struct {
		Email    string
		Password string
	}{}
	if err := c.Bind(str); err != nil {
		c.Flash().Add("error", "email or password missing")
		return c.Redirect(http.StatusUnauthorized, "/admin/login")
	}
	// check email & password
	usr := new(models.User)
	if err := models.DB.Where("email = ?", str.Email).First(usr); err != nil {
		c.Flash().Add("error", "email not found")
		return c.Redirect(http.StatusFound, "/admin/login")
	}
	if usr == nil {
		c.Flash().Add("error", "no such user")
		return c.Redirect(http.StatusFound, "/admin/login")
	}

	if !IsUser(usr, str.Password) {
		c.Flash().Add("error", "invalid password")
		return c.Redirect(http.StatusFound, "/admin/login")
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
		return c.Redirect(http.StatusFound, "/admin/login")
	}
	c.Session().Set("_token", token)
	c.Session().Save()

	return c.Redirect(http.StatusFound, "/login")
}
