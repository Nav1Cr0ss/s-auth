package handler

import (
	"context"

	oapi "github.com/Nav1Cr0ss/s-auth/internal/api"
	"github.com/Nav1Cr0ss/s-auth/internal/app"
	"github.com/Nav1Cr0ss/s-lib/logger"
)

type Handler struct {
	log *logger.Logger
	a   app.Application
}

func NewHandler(log *logger.Logger, a app.Application) Handler {
	h := Handler{
		log: log,
		a:   a,
	}

	return h
}

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (b Auth) HandleBearerAuth(ctx context.Context, operationName string, t oapi.BearerAuth) (context.Context, error) {
	return ctx, nil
}
