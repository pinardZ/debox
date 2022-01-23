package main

import (
	"fmt"
	"github.com/pinardZ/debox/apigateway/library"
	"log"
)

func main() {
	fmt.Println("hello debox")

	library.BSCClient = NewBSCCli()

	route := NewRouter()
	var err error
	if err = route.Run("0.0.0.0:8080"); err != nil {
		log.Print(err)
	}
}
