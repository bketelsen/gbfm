package gbfm

import (
	"strconv"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// homeHandler is a default handler to serve up a home page
func homeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("gbfm/index.html"))
}

func idFromSlug(s string) (uint, error) {
	pieces := strings.Split(s, "-")
	if len(pieces) < 1 {
		return 0, errors.New("bad article")
	}
	id := pieces[0]
	uid, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("bad article ID")
	}
	return uint(uid), nil
}
