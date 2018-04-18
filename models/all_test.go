package models

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type ModelSuite struct {
	*suite.Model
}

func TestModelSuite(t *testing.T) {
	ms := &ModelSuite{suite.NewModel()}
	suite.Run(t, ms)
}
