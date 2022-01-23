package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/pinardZ/debox/bscserver/proto"
	"github.com/pinardZ/debox/bscserver/service"
	"log"
)

func main() {
	server := micro.NewService(
		micro.Name("debox.service.bsc"),
		micro.Address(":8096"),
		//micro.Registry(etcd.NewRegistry(registry.Addrs(":12379"))),
	)
	server.Init()
	accountService := &service.AccountService{}

	var err error
	if err = proto.RegisterBSCServiceHandler(server.Server(), accountService); err != nil {
		log.Print(err)
		return
	}

	if err = server.Run(); err != nil {
		log.Print(err)
	}
}
