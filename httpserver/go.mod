module github.com/pinardZ/debox/apigateway

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.1
	github.com/micro/go-micro/v2 v2.9.1
    github.com/micro/go-plugins/registry/etcdv3/v2 v2.9.1
	google.golang.org/protobuf v1.24.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
