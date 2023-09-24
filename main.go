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

	// posts, err := queries.GetPostsByIds(ctx, []int32{1, 2, 14})
	// trataErro(err)
	// for _, post := range posts {
	// 	fmt.Println(post)
	// }

	// total, err := queries.CountPosts(ctx)
	// trataErro(err)
	// fmt.Printf("Existem %d posts no banco\n", total)

	report, err := queries.CountPostsByAuthor(ctx)
	trataErro(err)
	for _, r := range report {
		fmt.Printf("%s possui %d posts\n", r.Author.String, r.Count)
	}

}

func trataErro(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
