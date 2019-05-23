package handler

import (
	"net/http"
	"thing/platform/newsfeed"

	"github.com/qkgo/yin"
)

// PostsGet [GET] /posts
func PostsGet(feed newsfeed.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		items := feed.Get()
		res.SendJSON(items)
	}
}
