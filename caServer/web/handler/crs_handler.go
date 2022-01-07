package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/caServer/common/issue"
	"github.com/FuradWho/BlockchainDataColla/caServer/proto"
)

/*
const (
	PublicKeyFile  = "/home/fabric/GolandProjects/BlockchainDataColla/caServer/msp/signcert/ca.pem"
	PrivateKeyFile = "/home/fabric/GolandProjects/BlockchainDataColla/caServer/msp/keystore/ca.key"
)

*/

const (
	PublicKeyFile  = "E:\\projects\\BlockchainDataColla\\caServer\\msp\\signcert\\ca.pem"
	PrivateKeyFile = "E:\\projects\\BlockchainDataColla\\caServer\\msp\\keystore\\ca.key"
)

type CrsHandler struct {
	issue *issue.IssueCert
}

func (c *CrsHandler) setup() {

	cert := new(issue.IssueCert)

	cert.GetPublicKey(PublicKeyFile)
	cert.GetPrivateKey(PrivateKeyFile)

	c.issue = cert
}

func (c *CrsHandler) SendCsr(ctx context.Context, req *proto.CsrRequest, resp *proto.CsrResponse) error {
	csrDER := req.Csr
	cn := req.Cn

	if csrDER == nil {
		err := errors.New("failed to load empty crs")
		resp.Code = "400"
		resp.Msg = "failed to load empty crs!"

		return err
	}

	c.setup()
	err, crtBytes := c.issue.CrsCreateCrt(csrDER, cn)
	if err != nil {
		return err
	}

	resp.Code = "200"
	resp.Msg = "success to issue the cert !"
	resp.FileName = cn + ".crt"
	resp.FileSize = int32(len(crtBytes))
	resp.Crt = crtBytes

	fmt.Printf("%+v \n", resp)
	return nil

}

func (c *CrsHandler) GetCaCrt(ctx context.Context, req *proto.CaRequest, resp *proto.CaResponse) error {

	c.setup()
	err, crtBytes := c.issue.GetCaCrt()
	if err != nil {
		return err
	}
	resp.Code = "200"
	resp.Msg = "success to issue the cert !"
	resp.FileName = "ca.pem"
	resp.FileSize = int32(len(crtBytes))
	resp.CaCrt = crtBytes

	return nil
}
