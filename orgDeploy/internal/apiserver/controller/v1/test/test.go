package test

import (
	srvv1 "github.com/FuradWho/BlockchainDataColla/orgDeploy/internal/apiserver/service/v1"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/internal/apiserver/store"
)

type TestController struct {
	srv srvv1.Service
}

func NewTestController(store store.Factory) *TestController {
	return &TestController{
		srv: srvv1.NewService(store),
	}
}
