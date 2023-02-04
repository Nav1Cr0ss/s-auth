package handler

import (
	"context"

	oapi "github.com/Nav1Cr0ss/s-auth/internal/api"
	"github.com/Nav1Cr0ss/s-auth/internal/app"
	"github.com/Nav1Cr0ss/s-auth/pkg/logger"
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

type Bearer struct {
}

func NewBearer() *Bearer {
	return &Bearer{}
}

func (b Bearer) HandleBearerAuth(ctx context.Context, operationName string, t oapi.BearerAuth) (context.Context, error) {
	return ctx, nil
}
