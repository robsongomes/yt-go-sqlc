package main

import (
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/robsongomes/yt-go-sqlc/blog"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	slog.Info("Conectou com o banco")

	queries := blog.New(db)

	posts, err := queries.ListPosts(ctx)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	for _, post := range posts {
		slog.Info(post.Title)
	}
}
