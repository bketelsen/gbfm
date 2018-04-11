package content

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gophersnacks/gbfm/pkg/system/db"
)

// CacheControl sets the default cache policy on static asset responses
func CacheControl(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		res := c.Response()
		req := c.Request()
		cacheDisabled := db.ConfigCache("cache_disabled").(bool)
		if cacheDisabled {
			res.Header().Add("Cache-Control", "no-cache")
			return next(c)
		}
		age := int64(db.ConfigCache("cache_max_age").(float64))
		etag := db.ConfigCache("etag").(string)
		if age == 0 {
			age = db.DefaultMaxAge
		}
		policy := fmt.Sprintf("max-age=%d, public", age)
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
