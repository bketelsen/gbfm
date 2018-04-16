package content

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
)

const defaultMaxAge = 60

// CacheControl sets the default cache policy on static asset responses
func CacheControl(
	next buffalo.Handler,
	cacheDisabled bool,
	maxAge int,
	etag string,
) buffalo.Handler {
	return func(c buffalo.Context) error {
		res := c.Response()
		req := c.Request()
		if cacheDisabled {
			res.Header().Add("Cache-Control", "no-cache")
			return next(c)
		}
		if maxAge == 0 {
			maxAge = defaultMaxAge
		}
		policy := fmt.Sprintf("max-age=%d, public", maxAge)
		res.Header().Add("ETag", etag)
		res.Header().Add("Cache-Control", policy)

		if match := req.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				res.WriteHeader(http.StatusNotModified)
				return nil
			}
		}
		return next(c)
	}
}
