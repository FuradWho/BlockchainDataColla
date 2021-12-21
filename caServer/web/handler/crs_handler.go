package handler

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/caServer/proto"
)

type CrsHandler struct {
}

func (c *CrsHandler) CRS(ctx context.Context, req *proto.Request, resp *proto.Response) error {
	csrDER := req.Crs
	if csrDER == nil {
		return nil
	}

	csr, err := x509.ParseCertificateRequest(csrDER)
	if err != nil {
		return nil
	}

	fmt.Println(csr.PublicKey)

	resp.Code = "200"
	return nil
}

func (c *CrsHandler) GetTest(ctx context.Context, req *proto.Request, resp *proto.Response) error {
	fmt.Printf("%+s \n", req.Crs)
	resp.Code = "200"
	return nil
}
