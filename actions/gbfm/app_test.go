package gbfm

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/technosophos/moniker"
)

var namer = moniker.New()

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}
