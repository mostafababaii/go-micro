package grpc

import (
	"context"

	"github.com/mostafababaii/go-micro/services/auth/application"
	"github.com/mostafababaii/go-micro/services/auth/config"
	"github.com/mostafababaii/go-micro/services/auth/domain"
	"github.com/mostafababaii/go-micro/services/auth/infrastructure/persistence"
)

type GrpcServer struct {
	model *application.TokenInteractor
}

func NewGrpcServer() *GrpcServer {
	conn, err := config.NewDBConnection()
	if err != nil {
		panic(err)
	}
	return &GrpcServer{
		&application.TokenInteractor{
			Repository: persistence.NewTokenRepository(conn),
		},
	}
}

func (h *GrpcServer) ValidateToken(ctx context.Context, r *ValidateRequest) (*ValidateResponse, error) {
	token, err := h.model.ValidateToken(context.Background(), r.GetToken())
	if err != nil {
		panic(err)
	}
	return convertToGrpcValidate(token), err
}

func (h *GrpcServer) GetToken(ctx context.Context, r *TokenRequest) (*TokenResponse, error) {
	token, err := h.model.AddToken(context.Background(), uint(r.GetUser()))
	if err != nil {
		panic(err)
	}
	return convertToGrpcToken(token), err
}

func convertToGrpcToken(t *domain.Token) *TokenResponse {
	return &TokenResponse{
		Token: t.AccessToken,
	}
}

func convertToGrpcValidate(validate bool) *ValidateResponse {
	return &ValidateResponse{
		Status: validate,
	}
}
