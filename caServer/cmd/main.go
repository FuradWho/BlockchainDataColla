package main

import (
	proto "github.com/FuradWho/BlockchainDataColla/caServer/proto"
	_ "github.com/FuradWho/BlockchainDataColla/caServer/third_party/logger"
	"github.com/FuradWho/BlockchainDataColla/caServer/web/handler"
	consul "github.com/asim/go-micro/plugins/registry/consul/v3"
	micro "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"

	log "github.com/sirupsen/logrus"
)

const (
	ServerName = "FuradWho.BlockchainDataColla.caServer"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(
		registry.Addrs("192.168.2.4:8500"))
}

func main() {

	service := micro.NewService(
		micro.Name(ServerName),
		micro.Registry(consulReg),
		micro.Version("1.0"))

	service.Init()

	err := proto.RegisterCrsServiceHandler(service.Server(), new(handler.CrsHandler))
	if err != nil {
		log.Errorln(err)
	}

	//TODO

	_ = service.Run()
}
