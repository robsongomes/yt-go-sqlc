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

	//iniciar uma transação
	tx, err := db.BeginTx(ctx, nil)
	trataErro(err)

	qtx := queries.WithTx(tx)

	defer tx.Commit()

	post, err := qtx.GetPostById(ctx, 1)
	// tx.Rollback()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}

	err = qtx.SetPostViews(ctx, post.ID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}

	// err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

func trataErro(err error) {
	if err != nil {
		slog.Error(err.Error())
		// panic(err)
	}
}
