module github.com/pinardZ/debox/bscserver

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.8
	github.com/golang/protobuf v1.4.3
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	google.golang.org/protobuf v1.24.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
