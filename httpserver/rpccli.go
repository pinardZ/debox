package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/pinardZ/debox/apigateway/proto"
)

func NewBSCCli() proto.BSCService {
	client := micro.NewService(
		micro.Name("debox.client.bsc"),
		//micro.Registry(etcd.NewRegistry(registry.Addrs(":12379"))),
	)
	return proto.NewBSCService("debox.service.bsc", client.Client())
}
