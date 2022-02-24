package main

import (
	"context"
	"fmt"
	cert "github.com/FuradWho/BlockchainDataColla/orgDeploy/common/cert_apply"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/common/micro_services"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/proto/crs"
	fabric "github.com/FuradWho/BlockchainDataColla/orgDeploy/proto/fabric"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
	"github.com/asim/go-micro/v3"
	"github.com/prometheus/common/log"
	"time"
)

const (
	fabricServerName = "FuradWho.BlockchainDataColla.fabricDeploy"
	caServerName     = "FuradWho.BlockchainDataColla.caServer"
)

func Conn() {

	fabricOption, err := micro_services.NewFabricOption(func(option *micro_services.Option) {
		option.ServerName = fabricServerName
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

func CRT() {
	caOption, err := micro_services.NewCaOption(func(option *micro_services.Option) {
		option.ServerName = caServerName
	})
	if err != nil {
		log.Errorln(err)
	}

	microservice := micro.NewService(
		micro.Name(caOption.Option.ServerName),
		micro.Registry(caOption.Option.Registry),
	)

	microservice.Init()

	test := crs.NewCrsService(caOption.Option.ServerName, microservice.Client())

	certInfo := new(cert.Crt)
	err = certInfo.CreatePairKey()
	if err != nil {
		log.Errorln(err)
	}

	csrDER, err := certInfo.CreateCSR()
	if err != nil {
		log.Errorln(err)
	}

	resp, err := test.SendCsr(context.Background(), &crs.CsrRequest{
		Cn:  "node",
		Csr: csrDER,
	})
	if err != nil {
		log.Errorln(err)
	}

	err = certInfo.SaveCSR(resp.Crt)
	if err != nil {
		log.Errorln(err)
	}

}
func main() {
	Conn()
}
