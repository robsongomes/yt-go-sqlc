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

	params := blog.CreatePostAndReturnPostParams{
		Title:   "Como inserir com sqlc 4",
		Content: "Você pode usar o sql para inserir dados no banco",
		Slug:    "como_inserir_com_sqlc_4",
		Author:  sql.NullString{String: "robson", Valid: false},
	}

	// err = queries.CreatePost(ctx, params)
	// if err != nil {
	// 	panic(err)
	// }

	post, err := queries.CreatePostAndReturnPost(ctx, params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Post salvo: %v\n", post)
}
