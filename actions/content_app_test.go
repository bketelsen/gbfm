package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type ContentActionSuite struct {
	*suite.Action
}

func Test_Content(t *testing.T) {
	as := &ActionSuite{suite.NewAction(ContentApp())}
	suite.Run(t, as)
}
