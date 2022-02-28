package test

import (
	"encoding/json"
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/msg_client"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/model"
	"testing"
)

func TestServiceQuery(t *testing.T) {
	client := msg_client.FabricClient{}
	err := client.Init()
	if err != nil {
		fmt.Println(err)
	}

	serviceSetup := model.ServiceSetup{
		ChaincodeID: "msg_cc",
		Client:      client.ChannelClient,
	}

	result, err := serviceSetup.FindMsgInfoByMsgId("001")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var msg model.MsgBody
		json.Unmarshal(result, &msg)
		fmt.Println("根据 MsgId 查询信息成功：")
		fmt.Printf("%+v \n", msg)
	}

}

func TestServiceAdd(t *testing.T) {
	client := msg_client.FabricClient{}
	err := client.Init()
	if err != nil {
		fmt.Println(err)
	}

	serviceSetup := model.ServiceSetup{
		ChaincodeID: "msg_cc",
		Client:      client.ChannelClient,
	}
	msgBody := model.MsgBody{
		MsgId:     "001",
		Sender:    "zhansan",
		Recipient: "lisi",
		Body:      "1111111111111111",
		DataHash:  "1212",
		TimeDate:  "2022",
	}

	msg, err := serviceSetup.SaveMsg(msgBody)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
}
