package content

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/nilslice/jwt"
)

type ActionSuite struct {
	*suite.Action
}

func TestActions(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}

func (as *ActionSuite) login() error {
	claims := map[string]interface{}{}
	token, err := jwt.New(claims)
	if err != nil {
		return err
	}
	as.Session.Set("_token", token)
	if err := as.Session.Save(); err != nil {
		return err
	}
	return nil
}
