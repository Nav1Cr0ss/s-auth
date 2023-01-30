package app

import (
	"context"
	repository "github.com/Nav1Cr0ss/s-auth/internal/adapters/repository/sqlc"
)

func (a Application) GetAccountByLogin(ctx context.Context, phone string) (repository.User, error) {
	user, err := a.repo.GetAccountByLogin(ctx, phone)
	if err != nil {
		return user, err
	}
	return user, err

}
