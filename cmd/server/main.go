package main

import (
	db "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository"
	repository "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
	oapi "github.com/Nav1Cr0ss/s-auth/internal/api"
	"github.com/Nav1Cr0ss/s-auth/internal/app"
	"github.com/Nav1Cr0ss/s-auth/internal/config"
	handler "github.com/Nav1Cr0ss/s-auth/internal/ports/http"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewConfiguration,
			func(cfg *config.Config) *logger.Logger {
				return logger.NewLogger(cfg.Debug)
			},
			fx.Annotate(
				db.NewDB,
				fx.As(new(repository.DBTX)),
			),
			fx.Annotate(
				repository.New,
				fx.As(new(repository.Querier)),
			),

			app.NewApplication,
			fx.Annotate(
				handler.NewHandler,
				fx.As(new(oapi.Handler)),
			),
			fx.Annotate(
				handler.NewBearer,
				fx.As(new(oapi.SecurityHandler)),
			),
			//handler.NewHandler,
			oapi.NewServer,
		),
		fx.Invoke(
			handler.StartServer,
		),
	).Run()
}
