package main

import (
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/pkg/setting"
	fabric "github.com/FuradWho/BlockchainDataColla/orgDeploy/service/fabric_service"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
)

func main() {
	fabric.Conn()
}
