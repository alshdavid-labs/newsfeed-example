package handler

import (
	"net/http"
	"thing/platform/newsfeed"

	"github.com/qkgo/yin"
)

// PostsPostBody [POST] /posts
type PostsPostBody struct {
	Content string `json:"content"`
}

// PostsPost [POST] /posts
func PostsPost(feed newsfeed.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := PostsPostBody{}
		req.BindBody(&body)
		insert := newsfeed.Item{
			Content: body.Content,
		}
		items := feed.Set(insert)
		res.SendJSON(items)
	}
}
