package web

import (
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
)

// Translator is used by apps to translate text
var Translator *i18n.Translator

func init() {
	var err error
	if Translator, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
		panic(err)
	}
}
