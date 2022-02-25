package main

import (
	"fmt"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/pkg/setting"
	_ "github.com/FuradWho/BlockchainDataColla/orgDeploy/third_party/logger"
	"github.com/go-pg/pg/v10"
)

func main() {
	//r := routers.SetRouter()
	////启动端口为8085的项目
	//r.Run(":8081")

	opt, err := pg.ParseURL("postgres://furad:furad@localhost:5432/db_micro")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	fmt.Println(db.String())

}
