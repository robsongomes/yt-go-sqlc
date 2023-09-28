package main

import (
	"context"
	"database/sql"
	"fmt"
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

	report, err := queries.GetPostsViews(ctx)
	trataErro(err)

	fmt.Println("views --> post")
	for _, r := range report {
		fmt.Printf("%d --> %s\n", r.PostsView.Views, r.Post.Title)
	}
}

func trataErro(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
