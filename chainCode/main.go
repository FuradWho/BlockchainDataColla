package main

import (
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/chaincode/chaincode/edu_chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	err := shim.Start(new(edu_chaincode.EducationChaincode))
	if err != nil {
		fmt.Printf("启动EducationChaincode时发生错误: %s", err)
	}
}
