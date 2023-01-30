package app

import (
	"context"
	"github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
)

func (a Application) GetUser(ctx context.Context, id int64) (repository.User, error) {
	panic("implement me")
}
