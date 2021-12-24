package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	cert "github.com/FuradWho/BlockchainDataColla/orgDeploy/common/cert_apply"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/proto/crs"
	fabric "github.com/FuradWho/BlockchainDataColla/orgDeploy/proto/fabric"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
	grpc "github.com/asim/go-micro/plugins/client/grpc/v3"
	consul "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"io/ioutil"
	"time"
)

const (
	ServerName   = "FuradWho.BlockchainDataColla.fabricDeploy"
	caServerName = "FuradWho.BlockchainDataColla.caServer"
	serverCert   = "/home/fabric/GolandProjects/BlockchainDataColla/caServer/msp/signcert/ca.pem"
	clientKey    = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/keystore/client_private_key.pem"
	clientCert   = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/signcerts/client-ca-cert.crt"
)

func main() {
}

func setup() {
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
		RootCAs:            certPool,
		Certificates:       []tls.Certificate{x509KeyPair},
		InsecureSkipVerify: false,
	}

	//services, err := consulReg.ListServices()
	//if err != nil {
	//	return
	//}

	//grpcserver := grpc.NewServer()
	grpcserver := grpc.NewClient()
	//x509KeyPair, err := tls.LoadX509KeyPair("/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/signcerts/peer0.org1.example.com-cert.pem", "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/keystore/priv_sk")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//grpc.AuthTLS(&tls.Config{Certificates: []tls.Certificate{x509KeyPair}})

	grpcserver.Init(grpc.AuthTLS(tlsConfig))

	for {

		microservice := micro.NewService(
			micro.Client(grpcserver),
			micro.Name(ServerName),
			micro.Registry(consulReg),
		)

		microservice.Init()

		fabricClient := fabric.NewTestService(ServerName, microservice.Client())

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

func setup1() {

	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	microservice := micro.NewService(
		//		micro.Client(grpcserver),
		micro.Name(caServerName),
		micro.Registry(consulReg),
	)

	microservice.Init()

	test := crs.NewCrsService(caServerName, microservice.Client())

	certInfo := new(cert.Crt)
	err := certInfo.CreatePairKey()
	if err != nil {
		fmt.Println(err)
	}

	csrDER, err := certInfo.CreateCSR()
	if err != nil {
		fmt.Println(err)
	}

	resp, err := test.SendCsr(context.Background(), &crs.CsrRequest{
		Cn:  "node",
		Csr: csrDER,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = certInfo.SaveCSR(resp.Crt)
	if err != nil {
		fmt.Println(err)
	}

	//crt, err := test.GetCaCrt(context.Background(), &crs.CaRequest{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = ioutil.WriteFile("/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/ca/ca.pem", crt.CaCrt, 400)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
