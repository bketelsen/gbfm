package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type ContentActionSuite struct {
	*suite.Action
}

func Test_Content(t *testing.T) {
	app, close := ContentApp()
	defer close()
	as := &ActionSuite{suite.NewAction(app)}
	suite.Run(t, as)
}
