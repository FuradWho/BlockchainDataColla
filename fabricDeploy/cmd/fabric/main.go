package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	proto "github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto"
	handler "github.com/FuradWho/BlockchainDataColla/fabricDeploy/web/handler"
	nats "github.com/asim/go-micro/plugins/broker/nats/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	grpc "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/registry"
	"io/ioutil"
)

const (
	ServerName = "FuradWho.BlockchainDataColla.fabricDeploy"
	serverKey  = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/certprsk/keystore/server.key"
	serverCert = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/certprsk/signcert/server.crt"
	clientCert = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/signcerts/client.crt"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(
		registry.Addrs(":8500")) // 告知consul的端口号，如果走默认可以不填写
	// 冒号前面可填ip地址，默认localhost
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
		fmt.Println(err)
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
		RootCAs:            certPool,
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
