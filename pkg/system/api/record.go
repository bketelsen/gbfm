package api

import (
	"net/http"

	"github.com/gophersnacks/gbfm/pkg/system/api/analytics"
)

// Record wraps a HandlerFunc to record API requests for analytical purposes
func Record(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		go analytics.Record(req)

		next.ServeHTTP(res, req)
	})
}
