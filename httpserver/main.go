package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/pinardZ/debox/apigateway/controller"
	"github.com/pinardZ/debox/apigateway/library"
	"github.com/pinardZ/debox/apigateway/proto"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello debox")

	library.BSCClient = NewBSCCli()

	route := NewRouter()
	var err error
	if err = route.Run("0.0.0.0:80"); err != nil {
		log.Print(err)
	}
}

func NewRouter() *gin.Engine {
	route := gin.Default()

	registerAPI(route)

	return route
}

func registerAPI(route *gin.Engine) {
	route.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	routeGroup := route.Group("/api/")
	{
		routeGroup.GET("/:chain/account/balance", controller.AccountBalance)
		routeGroup.GET("/:chain/token/holderlist", controller.TokenHolderList)
	}
}

func NewBSCCli() proto.BSCService {
	client := micro.NewService(
		micro.Name("debox.client.bsc"),
	)
	// 初始化参数
	client.Init()
	return proto.NewBSCService("debox.service.bsc", client.Client())
}