package model

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	log "github.com/sirupsen/logrus"
	"time"
)

type MsgBody struct {
	ObjectType string        `json:"docType"`   // 消息类型
	MsgId      string        `json:"msg_id"`    // 信息 id
	Sender     string        `json:"sender"`    // 发送者
	Recipient  string        `json:"recipient"` // 接受者
	Body       string        `json:"body"`      // 消息体
	DataHash   string        `json:"hash"`      // 整体数据 hash
	TimeDate   string        `json:"time_date"` // 日期
	History    []HistoryItem // 当前 msg 的历史记录
}

type HistoryItem struct {
	TxId    string // 交易 id
	MsgBody MsgBody
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func registerEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		log.Errorf("注册链码事件失败: %s\n", err)
		return reg, notifier
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		log.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
