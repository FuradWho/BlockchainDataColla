package test

import (
	"github.com/FuradWho/BlockchainDataColla/caServer/common/issue"
	"testing"
)

func TestGetPublicKey(t *testing.T) {

	var issue issue.IssueCert

	issue.GetPublicKey("/home/fabric/GolandProjects/BlockchainDataColla/caServer/msp/signcert/ca.pem")
	issue.GetPrivateKey("/home/fabric/GolandProjects/BlockchainDataColla/caServer/msp/keystore/ca.key")
}
