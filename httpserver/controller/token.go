package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TokenHolderList(ctx *gin.Context) {
	chain := ctx.Param("chain")
	address := ctx.Query("address")

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": strings.Join([]string{"TokenHolderList", chain, address}, " "),
	})
}