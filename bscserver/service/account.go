package service

import (
	"context"
	"fmt"
	"github.com/pinardZ/debox/bscserver/proto"
)

type AccountService struct {
}

func (*AccountService) AccountBalance(c context.Context, req *proto.ReqAccountBalance,
	res *proto.RespAccountBalance) error {
	fmt.Println("hit here")
	res.Code = 0
	res.Message = "OK " + req.Address
	res.Result = "100"
	return nil
}
