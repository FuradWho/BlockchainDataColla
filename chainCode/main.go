package main

import (
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/chaincode/chaincode/message_chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	err := shim.Start(new(message_chaincode.MsgChaincode))
	if err != nil {
		fmt.Printf("main: An error occurred while starting EducationChaincode!: %s", err)
	}
}
