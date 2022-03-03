package micro_service

import (
	"github.com/asim/go-micro/v3"
	log "github.com/sirupsen/logrus"
)

const (
	ServerName   = "FuradWho.BlockchainDataColla.fabricDeploy"
	caServerName = "FuradWho.BlockchainDataColla.caServer"
)

var MicroOption Foo
var err error
var Service micro.Service

func init() {

	MicroOption, err = NewMicroOption(func(option *Option) {
		option.ServerName = ServerName
	})

	if err != nil {
		log.Errorln(err)
	}

	Service = micro.NewService(
		micro.Server(MicroOption.Option.Server),
		micro.Name(MicroOption.Option.ServerName),
		micro.Registry(MicroOption.Option.Registry),
		micro.Version(MicroOption.Option.Version),
		micro.Broker(MicroOption.Option.Broker),
	)

	Service.Init()
}
