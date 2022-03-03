package main

import (
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/micro_service"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/handler"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto/msg"
	_ "github.com/FuradWho/BlockchainDataColla/fabricDeploy/third_party/logger"
	log "github.com/sirupsen/logrus"
)

const (
	ServerName   = "FuradWho.BlockchainDataColla.fabricDeploy"
	caServerName = "FuradWho.BlockchainDataColla.caServer"
)

/*
func CRT() {

	caOption, err := micro_service.NewCaOption(func(option *micro_service.Option) {
		option.ServerName = caServerName
	})
	if err != nil {
		log.Errorln(err)
	}

	microservice := micro_service.NewService(
		micro_service.Name(caOption.Option.ServerName),
		micro_service.Registry(caOption.Option.Registry),
	)
	microservice.Init()

	test := csr.NewCrsService(caOption.Option.ServerName, microservice.Client())

	certInfo := new(cert.Crt)
	err = certInfo.CreatePairKey()
	if err != nil {
		fmt.Println(err)
	}

	csrDER, err := certInfo.CreateCSR()
	if err != nil {
		fmt.Println(err)
	}

	resp, err := test.SendCsr(context.Background(), &csr.CsrRequest{
		Cn:  "fabric",
		Csr: csrDER,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = certInfo.SaveCSR(resp.Crt)
	if err != nil {
		fmt.Println(err)
	}

	//crt, err := test.GetCaCrt(context.Background(), &csr.CaRequest{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = ioutil.WriteFile("/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/msp/ca/ca.pem", crt.CaCrt, 400)
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func Conn() {

	natsBroker := nats.NewBroker()
	natsBroker.Init(broker.Addrs("nats://192.168.175.129:4222"))
	natsBroker.Connect()

	fabricOption, err := micro_service.NewFabricOption(func(option *micro_service.Option) {
		option.Broker = natsBroker
		option.ServerName = ServerName
	})
	if err != nil {
		log.Errorln(err)
	}

	// grpcserver.Start()
	service := micro_service.NewService(
		micro_service.Server(fabricOption.Option.Server),
		micro_service.Name(fabricOption.Option.ServerName),
		micro_service.Registry(fabricOption.Option.Registry),
		micro_service.Version(fabricOption.Option.Version),
		micro_service.Broker(fabricOption.Option.Broker))

	service.Init()

	err = proto.RegisterTestServiceHandler(fabricOption.Option.Server, new(handler.TestService))

	if err != nil {
		fmt.Println(err)
	}

	_ = service.Run()
}

var ServiceSetup model.ServiceSetup

func Test() {
	client := msg_client.FabricClient{}
	err := client.Init()
	if err != nil {
		fmt.Println(err)
	}

	ServiceSetup = model.ServiceSetup{
		ChaincodeID: "msg_cc",
		Client:      client.ChannelClient,
	}
}

*/

func main() {

	option := micro_service.MicroOption.Option

	fmt.Println(option)
	err := msg.RegisterMsgServiceHandler(option.Server, new(handler.MsgService))
	if err != nil {
		log.Errorln(err)
	}

	err = micro_service.Service.Run()
	if err != nil {
		panic(err)
	}
}
