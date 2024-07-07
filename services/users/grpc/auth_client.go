package grpc

import (
	context "context"

	grpc "google.golang.org/grpc"
)

func GetToken(uid int32) (*TokenResponse, error) {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := NewAuthClient(conn)
	token, err := client.GetToken(context.Background(), &TokenRequest{User: uid})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidateToken(token string) (*ValidateResponse, error) {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := NewAuthClient(conn)
	validate, err := client.ValidateToken(context.Background(), &ValidateRequest{Token: token})
	if err != nil {
		return nil, err
	}
	return validate, nil
}
