package app

import (
	repository "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
	"github.com/Nav1Cr0ss/s-auth/pkg/logger"
)

type Application struct {
	repo repository.Querier
	//cfg     *config.Config
	log *logger.Logger
}

func NewApplication(repo repository.Querier, log *logger.Logger) Application {
	return Application{repo: repo, log: log}
}
