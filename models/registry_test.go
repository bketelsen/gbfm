package models

import (
	"github.com/jinzhu/inflection"
)

func (ms ModelSuite) TestRegistryFuncs() {
	r := ms.Require()
	for name, expectedFuncs := range registry {

		IDer, err := EmptyFromRegistry(name)
		r.NoError(err)
		r.Equal(expectedFuncs.empty(), IDer)

		IDer, err = EmptyFromRegistry(inflection.Plural(name))
		r.NoError(err)
		r.Equal(expectedFuncs.empty(), IDer)

		list, err := EmptyListFromRegistry(name)
		r.NoError(err)
		r.Equal(expectedFuncs.list(), list)

		list, err = EmptyListFromRegistry(inflection.Plural(name))
		r.NoError(err)
		r.Equal(expectedFuncs.list(), list)
	}
}
