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

	// err = queries.UpdatePostAuthor(ctx, sql.NullString{String: "miguel", Valid: true})
	// trataErro(err)
	// slog.Info("Registros atualizados")

	// err = queries.UpdatePostAuthorById(ctx, blog.UpdatePostAuthorByIdParams{
	// 	ID:     1,
	// 	Author: sql.NullString{String: "robson", Valid: true},
	// })
	// trataErro(err)
	// slog.Info("Registro atualizado")

	err = queries.DeletePostById(ctx, 16)
	trataErro(err)
	slog.Info("Registro deletado com sucesso")
}

func trataErro(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
