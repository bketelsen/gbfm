package snacks

import (
	"net/http"
)

func (as *ActionSuite) TestHome() {
	r := as.Require()
	res := as.HTML("/").Get()
	r.Equal(http.StatusOK, res.Code)
}
