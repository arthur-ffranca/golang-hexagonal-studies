package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	httpadapter "fastapi2/hexagonal/adapters/in/http"
	"fastapi2/hexagonal/adapters/out/postgres"
	"fastapi2/hexagonal/application/usecase"
)

func main() {
	ctx := context.Background()

	databaseURL := "postgres://app:app@127.0.0.1:5432/hexagonal?sslmode=disable"

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := postgres.NewUserRepository(db)
	createUser := usecase.NewCreateUserUseCase(repo)
	findUser := usecase.NewFindUserUseCase(repo)
	userHandler := httpadapter.NewUserHandler(createUser, findUser)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Create)
	mux.HandleFunc("/users/", userHandler.FindByID)

	log.Println("server running on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}