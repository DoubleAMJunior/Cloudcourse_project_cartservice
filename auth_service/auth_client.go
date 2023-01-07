package auth_service

import (
	"cart_service/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	con    *grpc.ClientConn
	Client AuthServiceClient
}

func (client *AuthClient) Close() {
	client.con.Close()
}

func GetAuthClient() (*AuthClient, error) {
	conn, err := grpc.Dial(config.AuthAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := NewAuthServiceClient(conn)
	result := &AuthClient{con: conn, Client: client}
	return result, nil
}
