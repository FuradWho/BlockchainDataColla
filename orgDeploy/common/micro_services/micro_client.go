package micro_services

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"
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
		return
	}
	option := MicroOption.Option

	Service = micro.NewService(
		micro.Client(option.Client),
		micro.Name(option.ServerName),
		micro.Registry(option.Registry),
		micro.Version(option.Version),
		micro.Broker(option.Broker),
	)

	Service.Init()

	//go func() {
	_, err := option.Broker.Subscribe("Msg", func(event broker.Event) error {
		fmt.Println(event.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	//}()
}

/*
	fabricOption, err := micro_services.NewFabricOption(func(option *micro_services.Option) {
		option.ServerName = setting.Conf.Service.FabricServerName
	})
	if err != nil {
		log.Errorln(err)
	}
	natsbroker := nats.NewBroker()
	natsbroker.Init(broker.Addrs("nats://192.168.175.129:4222"))
	err = natsbroker.Connect()
	if err != nil {
		fmt.Println(err)
	}

	microservice := micro.NewService(
		micro.Client(fabricOption.Option.Client),
		micro.Name(fabricOption.Option.ServerName),
		micro.Registry(fabricOption.Option.Registry),
		micro.Broker(natsbroker),
	)
	microservice.Init()

	msgClient := msg.NewMsgService(fabricOption.Option.ServerName, microservice.Client())
	response, err := msgClient.SaveMsgRpc(context.Background(), &msg.SaveMsgRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.GetMsg())
	subscribe, err := natsbroker.Subscribe("Msg", func(event broker.Event) error {
		fmt.Println(event.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
*/
