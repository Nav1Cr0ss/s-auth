// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetAccountByLogin(ctx context.Context, phone string) (User, error)
	GetUser(ctx context.Context, id int64) (User, error)
}

var _ Querier = (*Queries)(nil)