package content

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func TestActions(t *testing.T) {
	app, close := App()
	defer close()
	as := &ActionSuite{suite.NewAction(app)}
	suite.Run(t, as)
}
