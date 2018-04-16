package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type GBFMActionSuite struct {
	*suite.Action
}

func Test_GBFM(t *testing.T) {
	as := &ActionSuite{suite.NewAction(GBFMApp())}
	suite.Run(t, as)
}
