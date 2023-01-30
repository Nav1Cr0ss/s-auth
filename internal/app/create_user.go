package app

import (
	"context"
	repository "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
)

func (a Application) CreateUser(ctx context.Context, arg repository.CreateUserParams) (int64, error) {
	//TODO implement me
	panic("implement me")
}
