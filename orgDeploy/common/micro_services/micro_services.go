package micro_services

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/pkg/setting"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"io/ioutil"
)

/*
const (
	serverCert = "/home/furad/GolandProjects/BlockchainDataColla/caServer/msp/signcert/ca.pem"
	clientKey  = "/home/furad/GolandProjects/BlockchainDataColla/orgDeploy/msp/keystore/client_private_key.pem"
	clientCert = "/home/furad/GolandProjects/BlockchainDataColla/orgDeploy/msp/signcerts/client-ca-cert.crt"
)

*/

/*
const (
	serverCert = "E:\\projects\\BlockchainDataColla\\orgDeploy\\msp\\ca\\ca.pem"
	clientKey  = "E:\\projects\\BlockchainDataColla\\orgDeploy\\msp\\keystore\\client_private_key.pem"
	clientCert = "E:\\projects\\BlockchainDataColla\\orgDeploy\\msp\\signcerts\\client-ca-cert.crt"
)

*/

var consulReg = consul.NewRegistry(registry.Addrs(setting.Conf.Network.Ip + setting.Conf.Network.Port))

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

	x509KeyPair, err := tls.LoadX509KeyPair(setting.Conf.Path.ClientCert, setting.Conf.Path.ClientKey)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile(setting.Conf.Path.ServerCert)
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
