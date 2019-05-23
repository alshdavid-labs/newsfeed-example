package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"thing/httpd/handler"
	"thing/platform/newsfeed"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./newsfeed.db")
	feed := newsfeed.FromSQLite(db)
	
	// Alternatively use this
	// feed := newsfeed.InMemory()

	r := chi.NewRouter()
	r.Use(yin.SimpleLogger)

	r.Get("/posts", handler.PostsGet(feed))
	r.Post("/posts", handler.PostsPost(feed))

	fmt.Println("Watching on port: 3000")
	http.ListenAndServe(":3000", r)
}
