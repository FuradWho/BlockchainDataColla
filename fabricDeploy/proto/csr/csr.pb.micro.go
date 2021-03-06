// Code generated by protoc-gen-micro_service. DO NOT EDIT.
// source: csr.proto

// 指定等会文件生成出来的package

package csr

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CrsService service

func NewCrsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CrsService service

type CrsService interface {
	// 定义方法
	SendCsr(ctx context.Context, in *CsrRequest, opts ...client.CallOption) (*CsrResponse, error)
	GetCaCrt(ctx context.Context, in *CaRequest, opts ...client.CallOption) (*CaResponse, error)
}

type crsService struct {
	c    client.Client
	name string
}

func NewCrsService(name string, c client.Client) CrsService {
	return &crsService{
		c:    c,
		name: name,
	}
}

func (c *crsService) SendCsr(ctx context.Context, in *CsrRequest, opts ...client.CallOption) (*CsrResponse, error) {
	req := c.c.NewRequest(c.name, "CrsService.SendCsr", in)
	out := new(CsrResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crsService) GetCaCrt(ctx context.Context, in *CaRequest, opts ...client.CallOption) (*CaResponse, error) {
	req := c.c.NewRequest(c.name, "CrsService.GetCaCrt", in)
	out := new(CaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CrsService service

type CrsServiceHandler interface {
	// 定义方法
	SendCsr(context.Context, *CsrRequest, *CsrResponse) error
	GetCaCrt(context.Context, *CaRequest, *CaResponse) error
}

func RegisterCrsServiceHandler(s server.Server, hdlr CrsServiceHandler, opts ...server.HandlerOption) error {
	type crsService interface {
		SendCsr(ctx context.Context, in *CsrRequest, out *CsrResponse) error
		GetCaCrt(ctx context.Context, in *CaRequest, out *CaResponse) error
	}
	type CrsService struct {
		crsService
	}
	h := &crsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CrsService{h}, opts...))
}

type crsServiceHandler struct {
	CrsServiceHandler
}

func (h *crsServiceHandler) SendCsr(ctx context.Context, in *CsrRequest, out *CsrResponse) error {
	return h.CrsServiceHandler.SendCsr(ctx, in, out)
}

func (h *crsServiceHandler) GetCaCrt(ctx context.Context, in *CaRequest, out *CaResponse) error {
	return h.CrsServiceHandler.GetCaCrt(ctx, in, out)
}
