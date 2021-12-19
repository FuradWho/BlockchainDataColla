package main

import (
	proto "github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto"
	handler "github.com/FuradWho/BlockchainDataColla/fabricDeploy/web/handler"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

const (
	ServerName = "FuradWho.BlockchainDataColla.fabricDeploy"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(
		registry.Addrs(":8500")) // 告知consul的端口号，如果走默认可以不填写
	// 冒号前面可填ip地址，默认localhost
}

func main() {
	//ginRouter := gin.Default()
	//ginRouter.Handle("GET", "/test", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"data": "test success",
	//	})
	//})

	service := micro.NewService(
		micro.Address(":8581"),
		micro.Name(ServerName),
		micro.Registry(consulReg),
		micro.Version("1.0"))

	service.Init()

	err := proto.RegisterTestServiceHandler(service.Server(), new(handler.TestService))
	if err != nil {
		return
	}

	_ = service.Run()
}
