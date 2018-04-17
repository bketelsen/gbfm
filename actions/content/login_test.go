package content

import (
	"net/http"

	"github.com/nilslice/jwt"
)

func (as ActionSuite) TestShowLoginForm() {
	r := as.Require()

	// no session key exists
	res := as.HTML("/admin/login").Get()
	r.Equal(http.StatusOK, res.Code)

	// session key exists, but it's not a string
	as.Session.Set("_token", 123)
	r.NoError(as.Session.Save())
	res = as.HTML("/admin/login").Get()
	r.Equal(http.StatusOK, res.Code)

	// session key exists and is a string, but it's not a valid JWT
	as.Session.Set("_token", "invalidjwt")
	r.NoError(as.Session.Save())
	res = as.HTML("/admin/login").Get()
	r.Equal(http.StatusOK, res.Code)

	// session key exists and is a valid JWT token
	// create new token
	claims := map[string]interface{}{}
	token, err := jwt.New(claims)
	r.NoError(err)
	as.Session.Set("_token", token)
	r.NoError(as.Session.Save())
	res = as.HTML("/admin/login").Get()
	r.Equal(http.StatusFound, res.Code)
	r.Equal("/admin", res.Header().Get("Location"))

	// TODO: introspect some of the html. I (Aaron) have some code laying
	// around that does this. https://github.com/gophersnacks/gbfm/issues/19
}

func (as ActionSuite) TestTryLogin() {
	// TODO
	as.T().Log("TODO")
}
