package micro_services

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	grpc "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"io/ioutil"
)

const (
	//serverKey  = "E:\\projects\\BlockchainDataColla\\fabricDeploy\\msp\\keystore\\fabric_private_key.pem"
	//serverCert = "E:\\projects\\BlockchainDataColla\\fabricDeploy\\msp\\signcert\\client-ca-cert.crt"
	//clientCert = "E:\\projects\\BlockchainDataColla\\fabricDeploy\\msp\\ca\\ca.pem"

	serverKey  = "/home/furad/GolandProjects/BlockchainDataColla/fabricDeploy/msp/keystore/fabric_private_key.pem"
	serverCert = "/home/furad/GolandProjects/BlockchainDataColla/fabricDeploy/msp/signcert/client-ca-cert.crt"
	clientCert = "/home/furad/GolandProjects/BlockchainDataColla/fabricDeploy/msp/ca/ca.pem"
)

var consulReg = consul.NewRegistry(registry.Addrs("192.168.175.129:8500"))

type Foo struct {
	Option Option
}

type Option struct {
	Registry   registry.Registry
	Server     server.Server
	Broker     broker.Broker
	ServerName string
	Version    string
}

type ModOption func(option *Option)

func NewFabricOption(modOption ModOption) (*Foo, error) {

	grpcServer := grpc.NewServer()

	x509KeyPair, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		fmt.Println(err)
	}
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile(clientCert)
	if err != nil {
		return nil, err
	}

	certPool.AppendCertsFromPEM(certBytes)

	grpcServer.Init(grpc.AuthTLS(&tls.Config{
		Certificates:       []tls.Certificate{x509KeyPair},
		ClientCAs:          certPool,
		InsecureSkipVerify: false,
	}))

	option := Option{
		Server:     grpcServer,
		Registry:   consulReg,
		ServerName: "default",
		Version:    "1.0",
	}

	modOption(&option)

	return &Foo{
		Option: option,
	}, nil
}

func NewCaOption(modOption ModOption) (*Foo, error) {

	option := Option{
		Registry:   consulReg,
		ServerName: "default",
	}

	modOption(&option)

	return &Foo{
		Option: option,
	}, nil

}

func WithServer(server server.Server) ModOption {
	return func(option *Option) {
		option.Server = server
	}
}

func WithServerName(serverName string) ModOption {
	return func(option *Option) {
		option.ServerName = serverName
	}
}

func WithRegistry(registry registry.Registry) ModOption {
	return func(option *Option) {
		option.Registry = registry
	}
}
