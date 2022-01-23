package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pinardZ/debox/apigateway/controller"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	registerAPI(route)

	return route
}

func registerAPI(route *gin.Engine) {
	routeGroup := route.Group("/api/")
	{
		routeGroup.GET("/:chain/account/balance", controller.AccountBalance)
		routeGroup.GET("/:chain/token/holderlist", controller.TokenHolderList)
	}
}
