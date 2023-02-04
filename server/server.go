package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nav1Cr0ss/s-auth/config"
	"github.com/Nav1Cr0ss/s-auth/internal/api"
)

func StartServer(srv *api.Server, cfg *config.Config) {
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), srv); err != nil {
		log.Fatal(err)
	}
}
