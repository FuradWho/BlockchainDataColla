// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: Test.proto

// 指定等会文件生成出来的package

package proto

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

// Api Endpoints for TestService service

func NewTestServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TestService service

type TestService interface {
	// 定义方法
	GetTest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type testService struct {
	c    client.Client
	name string
}

func NewTestService(name string, c client.Client) TestService {
	return &testService{
		c:    c,
		name: name,
	}
}

func (c *testService) GetTest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "TestService.GetTest", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceHandler interface {
	// 定义方法
	GetTest(context.Context, *Request, *Response) error
}

func RegisterTestServiceHandler(s server.Server, hdlr TestServiceHandler, opts ...server.HandlerOption) error {
	type testService interface {
		GetTest(ctx context.Context, in *Request, out *Response) error
	}
	type TestService struct {
		testService
	}
	h := &testServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TestService{h}, opts...))
}

type testServiceHandler struct {
	TestServiceHandler
}

func (h *testServiceHandler) GetTest(ctx context.Context, in *Request, out *Response) error {
	return h.TestServiceHandler.GetTest(ctx, in, out)
}
