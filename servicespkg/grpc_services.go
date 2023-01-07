package servicespkg

import (
	"cart_service/auth_service"
	"cart_service/product_service"
	"context"
	"time"
)

type GrpcServices struct {
	auth    *auth_service.AuthClient
	product *product_service.ProductClient
}

func (service *GrpcServices) Init() error {
	var err error
	service.auth, err = auth_service.GetAuthClient()
	if err != nil {
		return err
	}
	service.product, err = product_service.GetProductClient()
	if err != nil {
		return err
	}
	return nil
}

func (service *GrpcServices) Close() {
	service.auth.Close()
	service.product.Close()
}

func (service *GrpcServices) GetUser(jwt string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	gUser, err := service.auth.Client.GetUser(ctx, &auth_service.JsonWebToken{Jwt: jwt})
	if err != nil {
		return nil, err
	}
	result := &User{Id: gUser.XId}
	return result, nil
}

func (service *GrpcServices) GetProduct(pid int32) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	gUser, err := service.product.Client.GetProduct(ctx, &product_service.PID{Id: pid})
	if err != nil {
		return nil, err
	}
	result := &Product{Title: gUser.Title, Count: gUser.Count}
	return result, nil
}

func (service *GrpcServices) HasAccess(method string, path string, jwt string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	mthd := auth_service.Resource_Method_value[method]
	result, err := service.auth.Client.HasAccess(ctx, &auth_service.Resource{Path: path, Jwt: jwt, Method: auth_service.Resource_Method(mthd)})
	if err != nil {
		return false, err
	}
	return result.HasAccess, nil
}
