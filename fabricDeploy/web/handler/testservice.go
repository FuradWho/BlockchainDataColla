package handler

import (
	"context"
	pb "github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type TestService struct {
}

func (r *TestService) GetTest(ctx context.Context, in *pb.Request, out *pb.Response) error {
	out.Errno = "300"
	log.Infof("%s %s \n", time.Now(), "Grpc Test !!!")
	return nil
}
