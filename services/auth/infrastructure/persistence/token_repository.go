package persistence

import (
	"context"
	"errors"

	_ "github.com/go-sql-driver/mysql" // driver
	"github.com/gocql/gocql"
	"github.com/mostafababaii/go-micro/services/auth/domain"
	"github.com/mostafababaii/go-micro/services/auth/domain/repository"
)

type tokenRepository struct {
	session *gocql.Session
}

func NewTokenRepository(session *gocql.Session) repository.TokenRepository {
	return &tokenRepository{session: session}
}

func (r *tokenRepository) Get(ctx context.Context, id uint) (*domain.Token, error) {
	iter := r.queryRow(ctx, "select * from tokens where id=? limit 1", id)
	t := &domain.Token{}
	iter.Scan(&t.ID, &t.AccessToken)
	if t.AccessToken == "" {
		return nil, errors.New("token not found")
	}
	return t, nil
}

func (r *tokenRepository) Save(ctx context.Context, t *domain.Token) error {
	err := r.session.Query(
		"insert into tokens (id, access_token, user_id, expires) values (?, ?, ?, ?);",
		t.ID,
		t.AccessToken,
		t.UserId,
		t.Expires,
	).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *tokenRepository) queryRow(ctx context.Context, q string, args ...interface{}) *gocql.Iter {
	return r.session.Query(q).Iter()
}
