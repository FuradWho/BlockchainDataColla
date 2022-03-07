package main

import (
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/common/micro_services"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/pkg/setting"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
	"github.com/FuradWho/BlockchainDataColla/orgDeploy/web/routers"
	log "github.com/sirupsen/logrus"
)

func main() {

	r := routers.SetRouter()
	//启动端口为8081的项目
	err := r.Run(":8081")
	if err != nil {
		log.Errorln(err)
		return
	}
}
