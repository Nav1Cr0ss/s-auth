package handler

import (
	"context"
	"github.com/Nav1Cr0ss/s-auth/internal/api"
)

func (h Handler) LoginPost(ctx context.Context, req api.OptLogin) (*api.Token, error) {
	var resp *api.Token
	user, err := h.a.GetAccountByLogin(ctx, req.Value.GetLogin())
	if err != nil {
		return resp, err
	}

	token := api.Token{RefreshToken: user.Phone}
	return &token, nil
}
