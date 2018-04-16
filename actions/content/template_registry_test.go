package content

import (
	"github.com/jinzhu/inflection"
)

func (as *ActionSuite) TestGetTemplateNames() {
	r := as.Require()
	for name, expectedTN := range templateRegistry {
		actualTN, err := getTemplateNames(name)
		r.NoError(err)
		r.Equal(*expectedTN, *actualTN)
		actualTN, err = getTemplateNames(inflection.Plural(name))
		r.NoError(err)
		r.Equal(*expectedTN, *actualTN)
	}
}
