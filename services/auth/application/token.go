package application

import (
	"context"

	"github.com/mostafababaii/go-micro/services/auth/domain"
	"github.com/mostafababaii/go-micro/services/auth/domain/repository"
)

type TokenInteractor struct {
	Repository repository.TokenRepository
}

func (i TokenInteractor) AddToken(ctx context.Context, id uint) (*domain.Token, error) {
	token, err := domain.NewToken(id)
	if err != nil {
		return nil, err
	}
	err = i.Repository.Save(ctx, token)
	return token, err
}

func (i TokenInteractor) ValidateToken(ctx context.Context, token string) (bool, error) {
	//return i.Repository.Get(ctx, id)
	return true, nil
}
