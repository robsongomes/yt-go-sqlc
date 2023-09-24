// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package blog

import (
	"database/sql"
)

type Post struct {
	ID      int64
	Title   string
	Content string
	Slug    string
	Author  sql.NullString
}

type PostsView struct {
	PostID int64
	Views  int32
}
