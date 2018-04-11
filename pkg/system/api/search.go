package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gophersnacks/gbfm/pkg/system/db"
	"github.com/gophersnacks/gbfm/pkg/system/item"
	"github.com/gophersnacks/gbfm/pkg/system/search"
)

func searchContentHandler(res http.ResponseWriter, req *http.Request) {
	qs := req.URL.Query()
	t := qs.Get("type")
	// type must be set, future version may compile multi-type result set
	if t == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	it, ok := item.Types[t]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if hide(res, req, it()) {
		return
	}

	q, err := url.QueryUnescape(qs.Get("q"))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// q must be set
	if q == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	count, err := strconv.Atoi(qs.Get("count")) // int: determines number of posts to return (10 default, -1 is all)
	if err != nil {
		if qs.Get("count") == "" {
			count = 10
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	offset, err := strconv.Atoi(qs.Get("offset")) // int: multiplier of count for pagination (0 default)
	if err != nil {
		if qs.Get("offset") == "" {
			offset = 0
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// execute search for query provided, if no index for type send 404
	matches, err := search.TypeQuery(t, q, count, offset)
	if err == search.ErrNoIndex {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("[search] Error:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// respond with json formatted results
	bb, err := db.ContentMulti(matches)
	if err != nil {
		log.Println("[search] Error:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if we have matches, push the first as its matched by relevance
	if len(bb) > 0 {
		push(res, req, it(), bb[0])
	}

	var result = []json.RawMessage{}
	for i := range bb {
		result = append(result, bb[i])
	}

	j, err := fmtJSON(result...)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err = omit(res, req, it(), j)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	sendData(res, req, j)
}
