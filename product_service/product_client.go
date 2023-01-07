package product_service

import (
	"cart_service/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var config_addr = "./config.json"

type ProductClient struct {
	con    *grpc.ClientConn
	Client ProductServiceClient
}

func (client *ProductClient) Close() {
	client.con.Close()
}

func GetProductClient() (*ProductClient, error) {
	conn, err := grpc.Dial(config.ProductAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := NewProductServiceClient(conn)
	result := &ProductClient{con: conn, Client: client}
	return result, nil
}
