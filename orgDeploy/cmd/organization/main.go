package main

import (
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/pkg/setting"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/web/routers"
)

func main() {
	r := routers.SetRouter()
	//启动端口为8085的项目
	r.Run(":8081")
}
