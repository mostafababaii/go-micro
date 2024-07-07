package repository

import (
	"context"

	"github.com/mostafababaii/go-micro/services/auth/domain"
)

type TokenRepository interface {
	Get(ctx context.Context, id uint) (*domain.Token, error)
	Save(ctx context.Context, user *domain.Token) error
}
