package main

import (
	"context"
	"fmt"
	cert "github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/cert_apply"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/micro_services"
	proto "github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto/csr"
	_ "github.com/FuradWho/BlockchainDataColla/fabricDeploy/third_party/logger"
	handler "github.com/FuradWho/BlockchainDataColla/fabricDeploy/web/handler"
	nats "github.com/asim/go-micro/plugins/broker/nats/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"
	log "github.com/sirupsen/logrus"
)

const (
	ServerName   = "FuradWho.BlockchainDataColla.fabricDeploy"
	caServerName = "FuradWho.BlockchainDataColla.caServer"
)

func man() {

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

func main() {

	natsBroker := nats.NewBroker()
	natsBroker.Init(broker.Addrs("nats://192.168.0.95:4222"))
	natsBroker.Connect()

	err := natsBroker.Publish("test", &broker.Message{
		Header: map[string]string{"type": "test"},
		Body:   []byte("test broker nats"),
	})
	if err != nil {
		log.Errorf("%s \n", err)
	}

	fabricOption, err := micro_services.NewFabricOption(func(option *micro_services.Option) {
		option.Broker = natsBroker
		option.ServerName = ServerName
	})
	if err != nil {
		log.Errorln(err)
	}

	// grpcserver.Start()
	service := micro.NewService(
		micro.Server(fabricOption.Option.Server),
		micro.Name(fabricOption.Option.ServerName),
		micro.Registry(fabricOption.Option.Registry),
		micro.Version(fabricOption.Option.Version),
		micro.Broker(fabricOption.Option.Broker))

	service.Init()

	err = proto.RegisterTestServiceHandler(fabricOption.Option.Server, new(handler.TestService))

	if err != nil {
		fmt.Println(err)
	}

	_ = service.Run()
}
