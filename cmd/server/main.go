package main

import (
	"context"
	"database/sql"
	"github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
	api2 "github.com/Nav1Cr0ss/s-auth/internal/api"
	"github.com/Nav1Cr0ss/s-auth/internal/app"
	"github.com/Nav1Cr0ss/s-auth/internal/ports/http"
	"github.com/Nav1Cr0ss/s-auth/pkg/logger"
	_ "github.com/lib/pq"
	"net/http"
)

type Bearer struct {
}

func (b Bearer) HandleBearerAuth(ctx context.Context, operationName string, t api2.BearerAuth) (context.Context, error) {
	return ctx, nil
}

func main() {
	log := logger.NewLogger()

	db, err := sql.Open("postgres", "postgresql://nav1cr0ss:0608@localhost:5436/s-auth?sslmode=disable")
	if err != nil {
		return
	}

	repo := repository.New(db)

	application := app.NewApplication(repo, log)

	h := handler.NewHandler(log, application)

	srv, err := api2.NewServer(h, Bearer{})
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
