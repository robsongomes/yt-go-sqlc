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
	post, err := queries.UpdateContentOrAuthor(ctx, blog.UpdateContentOrAuthorParams{
		//Content: sql.NullString{String: "Conteúdo alterado", Valid: true},
		Author: sql.NullString{String: "robson", Valid: true},
		ID:     2,
	})
	trataErro(err)
	fmt.Println(post)
}

func trataErro(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
