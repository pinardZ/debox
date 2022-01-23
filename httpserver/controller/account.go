package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pinardZ/debox/apigateway/library"
	"github.com/pinardZ/debox/apigateway/proto"
	"net/http"
)

func AccountBalance(ctx *gin.Context) {
	chain := ctx.Param("chain")
	address := ctx.Query("address")

	fmt.Println("start " + chain + " " + address)

	resp, err := library.BSCClient.AccountBalance(context.TODO(), &proto.ReqAccountBalance{Address: address})
	if err != nil {
		fmt.Println("rpccli.bsc ral error, ", err.Error())
		return
	}

	fmt.Println("222 " + chain + " " + address)

	resp.Result += " " + chain

	ctx.JSON(http.StatusOK, resp)
}
