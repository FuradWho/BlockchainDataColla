package cert_service

import (
	"context"
	cert "github.com/FuradWho/BlockchainDataColla/orgDeploy/common/cert_apply"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/common/micro_services"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/pkg/setting"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/proto/crs"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
	"github.com/asim/go-micro/v3"
	"github.com/prometheus/common/log"
)

/*
const (
	caServerName = "FuradWho.BlockchainDataColla.caServer"
)
*/

func CRT() {
	caOption, err := micro_services.NewCaOption(func(option *micro_services.Option) {
		option.ServerName = setting.Conf.Service.CaServerName
	})
	if err != nil {
		log.Errorln(err)
	}

	microservice := micro.NewService(
		micro.Name(caOption.Option.ServerName),
		micro.Registry(caOption.Option.Registry),
	)

	microservice.Init()

	test := crs.NewCrsService(caOption.Option.ServerName, microservice.Client())

	certInfo := new(cert.Crt)
	err = certInfo.CreatePairKey()
	if err != nil {
		log.Errorln(err)
	}

	csrDER, err := certInfo.CreateCSR()
	if err != nil {
		log.Errorln(err)
	}

	resp, err := test.SendCsr(context.Background(), &crs.CsrRequest{
		Cn:  "node",
		Csr: csrDER,
	})
	if err != nil {
		log.Errorln(err)
	}

	err = certInfo.SaveCSR(resp.Crt)
	if err != nil {
		log.Errorln(err)
	}

}
