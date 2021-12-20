package handler

import (
	"context"
	"fmt"
	pb "github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto"
	"time"
)

type TestService struct {
}

func (r *TestService) GetTest(ctx context.Context, in *pb.Request, out *pb.Response) error {
	out.Errno = "300"
	fmt.Println(time.Now())
	fmt.Println("Grpc Test !!!")
	return nil
}
