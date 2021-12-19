package main

import (
	"context"
	"fmt"
	proto "github.com/FuradWho/BlockchainDataColla/orgDeploy/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"io/ioutil"
	"net/http"
)

const (
	ServerName = "FuradWho.BlockchainDataColla.fabricDeploy"
)

func callApi(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	services, err := consulReg.ListServices()
	if err != nil {
		return
	}

	for _, service := range services {
		microservice := micro.NewService(
			micro.Name(service.Name),
			micro.Registry(consulReg))

		microservice.Init()

		test := proto.NewTestService(ServerName, microservice.Client())

		getTest, err := test.GetTest(context.TODO(), &proto.Request{})
		if err != nil {
			return
		}
		fmt.Println(service.Name)
		fmt.Println(getTest)

	}

}
