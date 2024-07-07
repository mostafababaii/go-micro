package services

import "github.com/mostafababaii/go-micro/services/users/grpc"

var (
	AuthService = &authService{}
)

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

type authRequest struct {
	Username string
	Password string
}

func NewAuthRequest() *authRequest {
	return &authRequest{}
}

func (*authService) GetToken(userID uint) (string, error) {
	token, err := grpc.GetToken(int32(userID))
	if err != nil {
		return "", err
	}
	return token.Token, nil
}
