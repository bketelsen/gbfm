package web

import (
	"github.com/gobuffalo/buffalo"
)

// LayoutMiddleware is middleware to make it easy to use the common layout
// in templates/common_layout.html.
//
// common_layout.html is a skeleton layout for each property and is customized
// with various template variables. This middleware sets them
func LayoutMiddleware(title, navPartial, footerPartial string) buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			c.Set("title", title)
			c.Set("navPartial", navPartial)
			c.Set("footerPartial", footerPartial)
			return next(c)
		}
	}
}
