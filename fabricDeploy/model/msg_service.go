package model

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	log "github.com/sirupsen/logrus"
)

func (t *ServiceSetup) SaveMsg(edu MsgBody) (string, error) {

	eventID := "eventAddMsg"
	reg, notifier := registerEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(edu)
	if err != nil {
		log.Errorf("指定的edu对象序列化时发生错误: %s \n", err)
		return "", errors.New("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "addMsg",
		Args:        [][]byte{b, []byte(eventID)},
	}

	response, err := t.Client.Execute(req)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	return string(response.TransactionID), nil
}

func (t *ServiceSetup) FindMsgInfoByMsgId(MsgId string) ([]byte, error) {

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "queryMsgByID",
		Args:        [][]byte{[]byte(MsgId)},
	}
	response, err := t.Client.Query(req)
	if err != nil {
		log.Errorln(err)
		return []byte{0x00}, err
	}

	return response.Payload, nil
}

func (t *ServiceSetup) ModifyEdu(edu MsgBody) (string, error) {

	eventID := "eventModifyMsg"
	reg, notifier := registerEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(edu)
	if err != nil {
		log.Errorf("指定的edu对象序列化时发生错误: %s \n", err)
		return "", errors.New("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "updateMsg",
		Args:        [][]byte{b, []byte(eventID)},
	}
	response, err := t.Client.Execute(req)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	return string(response.TransactionID), nil
}
