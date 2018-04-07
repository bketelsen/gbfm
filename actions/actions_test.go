package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}

func Test_GBFM_ActionSuite(t *testing.T) {
	as := &gbfm.ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}

func Test_Snacks_ActionSuite(t *testing.T) {
	as := &snacks.ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}
