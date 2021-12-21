package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	cert "github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/cert_apply"
	proto "github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto/csr"
	_ "github.com/FuradWho/BlockchainDataColla/fabricDeploy/third_party/logger"
	handler "github.com/FuradWho/BlockchainDataColla/fabricDeploy/web/handler"
	nats "github.com/asim/go-micro/plugins/broker/nats/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	grpc "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/registry"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

const (
	ServerName   = "FuradWho.BlockchainDataColla.fabricDeploy"
	serverKey    = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/msp/keystore/fabric_private_key.pem"
	serverCert   = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/msp/signcert/client-ca-cert.crt"
	clientCert   = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/msp/ca/ca.pem"
	caServerName = "FuradWho.BlockchainDataColla.caServer"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(
		registry.Addrs(":8500")) // 告知consul的端口号，如果走默认可以不填写
	// 冒号前面可填ip地址，默认localhost
}

func e() {

	microservice := micro.NewService(
		//		micro.Client(grpcserver),
		micro.Name(caServerName),
		micro.Registry(consulReg),
	)

	microservice.Init()

	test := csr.NewCrsService(caServerName, microservice.Client())

	certInfo := new(cert.Crt)
	err := certInfo.CreatePairKey()
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

	natsbroker := nats.NewBroker()
	natsbroker.Init(broker.Addrs("nats://0.0.0.0:4222"))
	natsbroker.Connect()

	err := natsbroker.Publish("test", &broker.Message{
		Header: map[string]string{"type": "test"},
		Body:   []byte("test broker nats"),
	})
	if err != nil {
		log.Errorf("%s \n", err)
	}

	grpcserver := grpc.NewServer()

	x509KeyPair, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		fmt.Println(err)
	}
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile(clientCert)
	if err != nil {
		return
	}

	certPool.AppendCertsFromPEM(certBytes)

	grpcserver.Init(grpc.AuthTLS(&tls.Config{
		Certificates:       []tls.Certificate{x509KeyPair},
		ClientCAs:          certPool,
		InsecureSkipVerify: false,
	}))

	// grpcserver.Start()
	service := micro.NewService(
		micro.Server(grpcserver),
		micro.Name(ServerName),
		micro.Registry(consulReg),
		micro.Version("1.0"),
		micro.Broker(natsbroker))

	service.Init()

	err = proto.RegisterTestServiceHandler(grpcserver, new(handler.TestService))
	if err != nil {
		fmt.Println(err)
	}

	_ = service.Run()
}
