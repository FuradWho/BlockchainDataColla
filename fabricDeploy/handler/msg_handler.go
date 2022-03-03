package handler

import (
	"context"
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/msg_client"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/nats_service"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/model"
	pb "github.com/FuradWho/BlockchainDataColla/fabricDeploy/proto/msg"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

var serviceSetup model.ServiceSetup

func init() {
	client := msg_client.FabricClient{}
	err := client.Init()
	if err != nil {
		fmt.Println(err)
	}

	serviceSetup = model.ServiceSetup{
		ChaincodeID: "msg_cc",
		Client:      client.ChannelClient,
	}
}

type MsgService struct {
}

func (m *MsgService) SaveMsgRpc(ctx context.Context, in *pb.SaveMsgRequest, out *pb.SaveMsgResponse) error {

	log.Infoln("%s %s \n", time.Now(), "SaveMsgRpc !!!")

	var msg model.MsgBody
	v4, err := uuid.NewV4()
	if err != nil {
		return err
	}

	msg.MsgId = v4.String()
	msg.Sender = in.Sender
	msg.Recipient = in.Recipient
	msg.Body = in.Body
	msg.TimeDate = in.TimeDate
	msg.DataHash = in.DataHash

	transactionID, err := serviceSetup.SaveMsg(msg)
	if err != nil {
		return err
	}

	fmt.Println(transactionID)

	out.TxId = transactionID
	out.Msg = msg.MsgId
	out.Code = "200"

	nats_service.SaveMsg(out.Msg)

	return nil
}
