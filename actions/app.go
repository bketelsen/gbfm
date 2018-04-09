package actions

import (
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var translator *i18n.Translator

func init() {
	var err error
	if translator, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
		panic(err)
	}
}
