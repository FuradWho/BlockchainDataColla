package micro_services

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"io/ioutil"
)

const (
	serverCert = "/home/fabric/GolandProjects/BlockchainDataColla/caServer/msp/signcert/ca.pem"
	clientKey  = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/keystore/client_private_key.pem"
	clientCert = "/home/fabric/GolandProjects/BlockchainDataColla/orgDeploy/msp/signcerts/client-ca-cert.crt"
)

var consulReg = consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

type Foo struct {
	Option Option
}

type Option struct {
	Registry   registry.Registry
	Client     client.Client
	ServerName string
}

type ModOption func(option *Option)

func NewFabricOption(modOption ModOption) (*Foo, error) {

	x509KeyPair, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile(serverCert)
	if err != nil {
		return nil, err
	}

	certPool.AppendCertsFromPEM(certBytes)

	tlsConfig := &tls.Config{
		RootCAs:            certPool,
		Certificates:       []tls.Certificate{x509KeyPair},
		InsecureSkipVerify: false,
	}

	grpcServer := grpc.NewClient()
	err = grpcServer.Init(grpc.AuthTLS(tlsConfig))
	if err != nil {
		return nil, err
	}

	option := Option{
		Client:     grpcServer,
		Registry:   consulReg,
		ServerName: "default",
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

func WithClient(client client.Client) ModOption {
	return func(option *Option) {
		option.Client = client
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