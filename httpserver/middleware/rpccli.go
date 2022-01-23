package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pinardZ/debox/apigateway/proto"
)

func InitBSCRPCCliMiddleware(bscService proto.BSCService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = map[string]interface{}{
			"rpccli.bsc": bscService,
		}
		ctx.Next()
	}
}
