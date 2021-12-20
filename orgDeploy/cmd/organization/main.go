package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/proto"
	grpc "github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"io/ioutil"
	"time"
)

const (
	ServerName = "FuradWho.BlockchainDataColla.fabricDeploy"
	serverCert = "/home/fabric/GolandProjects/BlockchainDataColla/fabricDeploy/certprsk/signcert/server.crt"
	clientKey  = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/keystore/client.key"
	clientCert = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/signcerts/client.crt"
)

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	x509KeyPair, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		return
	}
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile(serverCert)
	if err != nil {
		return
	}

	certPool.AppendCertsFromPEM(certBytes)

	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{x509KeyPair},
	}

	//services, err := consulReg.ListServices()
	//if err != nil {
	//	return
	//}

	//grpcserver := grpc.NewServer()

	for {
		grpcserver := grpc.NewClient()
		//x509KeyPair, err := tls.LoadX509KeyPair("/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/signcerts/peer0.org1.example.com-cert.pem", "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/keystore/priv_sk")
		//if err != nil {
		//	fmt.Println(err)
		//}
		//grpc.AuthTLS(&tls.Config{Certificates: []tls.Certificate{x509KeyPair}})

		grpcserver.Init(grpc.AuthTLS(tlsConfig))
		microservice := micro.NewService(
			micro.Client(grpcserver),
			micro.Name(ServerName),
			micro.Registry(consulReg),
		)

		microservice.Init()

		test := proto.NewTestService(ServerName, microservice.Client())

		getTest, err := test.GetTest(context.TODO(), &proto.Request{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ServerName)
		fmt.Printf("%+v \n", getTest)
		time.Sleep(time.Second * 2)
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
