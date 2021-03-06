package fabric_service

import (
	"context"
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/common/micro_services"
	msg "github.com/FuradWho/BlockchainDataColla/orgDeploy/proto/msg"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
)

/*
const (
	fabricServerName = "FuradWho.BlockchainDataColla.fabricDeploy"
)

*/
/*
func Conn() {

	fabricOption, err := micro_services.NewFabricOption(func(option *micro_services.Option) {
		option.ServerName = setting.Conf.Service.FabricServerName
	})
	if err != nil {
		log.Errorln(err)
	}

	microservice := micro.NewService(
		micro.Client(fabricOption.Option.Client),
		micro.Name(fabricOption.Option.ServerName),
		micro.Registry(fabricOption.Option.Registry))
	microservice.Init()

	for {
		fabricClient := fabric.NewTestService(fabricOption.Option.ServerName, microservice.Client())
		response, err := fabricClient.GetTest(context.Background(), &fabric.Request{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response.GetErrno())
		time.Sleep(time.Second * 5)
	}

	//for {
	//	time.Sleep(time.Second * 5)
	//	natsbroker := nats.NewBroker()
	//	natsbroker.Init(broker.Addrs("nats://0.0.0.0:4222"))
	//	err := natsbroker.Connect()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("!!!")
	//	subscribe, err := natsbroker.Subscribe("test", func(event broker.Event) error {
	//		fmt.Println("11111")
	//		fmt.Println(event.Message().Body)
	//		fmt.Println(event.Message().Header)
	//		return nil
	//	})
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("???")
	//	fmt.Println(subscribe.Topic())
	//}

}

*/

func Msg() {

	option := micro_services.MicroOption.Option
	msgClient := msg.NewMsgService(option.ServerName, micro_services.Service.Client())
	response, err := msgClient.SaveMsgRpc(context.Background(), &msg.SaveMsgRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.GetMsg())
}
