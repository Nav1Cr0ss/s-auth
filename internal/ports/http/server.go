package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Nav1Cr0ss/s-auth/internal/api"
	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, srv *api.Server) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := http.ListenAndServe(":8080", srv); err != nil {
					log.Fatal(err)
				}

				return nil
			},
			//OnStop: func(ctx context.Context) error {
			//	err := srv.Echo.Shutdown(ctx)
			//	if err != nil {
			//		srv.log.Info().Err(err).Msg("couldn't stop server")
			//	}
			//
			//	return nil
			//},
		},
	)
}
