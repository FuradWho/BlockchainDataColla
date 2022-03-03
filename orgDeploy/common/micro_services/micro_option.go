package micro_services

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/pkg/setting"
	"github.com/asim/go-micro/plugins/broker/nats/v3"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"io/ioutil"
)

var consulReg = consul.NewRegistry(registry.Addrs(setting.Conf.Network.Ip))

type Foo struct {
	Option Option
}

type Option struct {
	Registry   registry.Registry
	Client     client.Client
	Broker     broker.Broker
	ServerName string
	Version    string
}

type ModOption func(option *Option)

func NewMicroOption(modOption ModOption) (Foo, error) {

	x509KeyPair, err := tls.LoadX509KeyPair(setting.Conf.Path.ClientCert, setting.Conf.Path.ClientKey)
	if err != nil {
		return Foo{}, err
	}
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile(setting.Conf.Path.ServerCert)
	if err != nil {
		return Foo{}, err
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
		return Foo{}, err
	}

	// nats
	natsBroker := nats.NewBroker()
	err = natsBroker.Init(broker.Addrs("nats://192.168.175.129:4222"))
	if err != nil {
		return Foo{}, err
	}
	err = natsBroker.Connect()
	if err != nil {
		return Foo{}, err
	}

	option := Option{
		Client:     grpcServer,
		Registry:   consulReg,
		Broker:     natsBroker,
		ServerName: "default",
		Version:    "latest",
	}

	modOption(&option)

	return Foo{
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
