package handler

import (
	repository "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
	"github.com/Nav1Cr0ss/s-auth/pkg/logger"
)

type Handler struct {
	log *logger.Logger
	a   repository.Querier
}

func NewHandler(log *logger.Logger, a repository.Querier) Handler {
	h := Handler{
		log: log,
		a:   a,
	}

	return h
}
