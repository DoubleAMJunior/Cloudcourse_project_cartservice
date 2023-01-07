package main

import (
	"cart_service/config"
	"cart_service/controler"
	middleWare "cart_service/middleware"
	"cart_service/servicespkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Initializing services")
	helperService := &servicespkg.GrpcServices{}
	err := helperService.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer helperService.Close()
	addcart := controler.AddCartController{Helper: helperService}
	getcart := controler.GetCartController{Helper: helperService}
	modifycart := controler.ModifyCartController{Helper: helperService}
	authMiddle := middleWare.UserAuthMiddleware{Helper: helperService}
	fmt.Println("Starting server")
	server := gin.New()
	server.Use(authMiddle.Action)
	server.POST("/add", addcart.AddCart)
	server.GET("/get", getcart.GetCart)
	server.POST("/modify", modifycart.ModifyCart)
	server.Run(config.ServerAddr)
}
