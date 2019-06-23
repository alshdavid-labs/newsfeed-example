package newsfeed

import (
	"database/sql"
)

// SQLite holds the DB connection
type SQLite struct {
	DB *sql.DB
}

// Get gets items from the database
func (s *SQLite) Get() []Item {
	items := []Item{}
	rows, _ := s.DB.Query("SELECT * FROM newsfeed_items")
	defer rows.Close()
	var id string
	var content string
	for rows.Next() {
		rows.Scan(&id, &content)
		items = append(items, Item{content})
	}
	return items
}

// Set stores items to the database
func (s *SQLite) Set(item Item) bool {
	stmt, _ := s.DB.Prepare("INSERT INTO newsfeed_items(content) values (?)")
	stmt.Exec(item.Content)
	return true
}

// FromSQLite creates a newfeed that uses sqlite
func FromSQLite(conn *sql.DB) *SQLite {
	stmt, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS
		newsfeed_items (
			ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			content	TEXT
		);
	`)
	stmt.Exec()
	return &SQLite{
		DB: conn,
	}
}
