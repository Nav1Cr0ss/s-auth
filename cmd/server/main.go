package main

import (
	"context"

	db "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository"
	repository "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
	oapi "github.com/Nav1Cr0ss/s-auth/internal/api"
	"github.com/Nav1Cr0ss/s-auth/internal/app"
	"github.com/Nav1Cr0ss/s-auth/internal/config"
	handler "github.com/Nav1Cr0ss/s-auth/internal/ports/http"
	"github.com/Nav1Cr0ss/s-auth/pkg/logger"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func main() {
	//log := logger.NewLogger()
	//
	//conf, err := app.NewConfiguration(log)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:0608@localhost:5436/s-auth?sslmode=disable", conf.Db.UserName))
	//if err != nil {
	//	return
	//}
	//
	//repo := repository.New(db)
	//
	//application := app.NewApplication(repo, log)
	//
	//h := handler.NewHandler(log, application)
	//
	//srv, err := api2.NewServer(h, Bearer{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err := http.ListenAndServe(":8080", srv); err != nil {
	//	log.Fatal(err)
	//}

	fx.New(
		fx.Provide(
			logger.NewLogger,
			config.NewConfiguration,

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
			func(lc fx.Lifecycle, log *logger.Logger, cfg *config.Config) {
				lc.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							log.Info("failed to fetch URL")

							return nil
						},
					},
				)
			},
		),
	).Run()
}
