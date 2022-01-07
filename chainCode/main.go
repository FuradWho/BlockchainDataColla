package main

import (
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/chaincode/chaincode/eduChaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	err := shim.Start(new(eduChaincode.EducationChaincode))
	if err != nil {
		fmt.Printf("启动EducationChaincode时发生错误: %s", err)
	}
}
