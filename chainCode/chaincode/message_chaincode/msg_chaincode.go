package message_chaincode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/chaincode/model/message_body"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const (
	DocType = "msgObj"
)

type MsgChaincode struct {
}

func (e *MsgChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
func (e *MsgChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	fun, args := stub.GetFunctionAndParameters()

	if fun == "addMsg" {
		return e.addMsg(stub, args) // 添加信息
	} else if fun == "queryMsgByID" {
		return e.queryMsgByID(stub, args) // 根据身份证号码及姓名查询详情
	} else if fun == "updateMsg" {
		return e.updateMsg(stub, args) // 根据证书编号更新信息
	} else if fun == "delMsg" {
		return e.delMsg(stub, args) // 根据证书编号删除信息
	}

	return shim.Error("Invoke: the specified function name is wrong!")

}

func PutMsg(stub shim.ChaincodeStubInterface, msg message_body.MsgBody) ([]byte, bool) {

	msg.ObjectType = DocType

	marshal, err := json.Marshal(msg)
	if err != nil {
		return nil, false
	}

	err = stub.PutState(msg.MsgId, marshal)
	if err != nil {
		return nil, false
	}

	return marshal, true

}

func GetMsg(stub shim.ChaincodeStubInterface, msgId string) (message_body.MsgBody, bool) {

	var msg message_body.MsgBody

	b, err := stub.GetState(msgId)
	if err != nil {
		return msg, false
	}

	if b == nil {
		return msg, false
	}

	err = json.Unmarshal(b, &msg)
	if err != nil {
		return msg, false
	}

	return msg, true
}

// 根据指定的查询字符串实现富查询
func getEduByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

func (e *MsgChaincode) addMsg(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("addMsg: input parameter error!")
	}

	var msg message_body.MsgBody

	err := json.Unmarshal([]byte(args[0]), &msg)
	if err != nil {
		return shim.Error("addMsg: unmarshal has error!")
	}

	_, exist := GetMsg(stub, msg.MsgId)
	if exist {
		return shim.Error("addMsg: msg has exist!")
	}

	_, saveFlag := PutMsg(stub, msg)
	if !saveFlag {
		return shim.Error("addMsg: save msg has error!")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("addMsg: success to save msg !"))
}

// 溯源
// args: entityID
func (e *MsgChaincode) queryMsgByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("queryMsgByID: input parameter error!")
	}

	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("queryMsgByID: query based on id failed!")
	}

	if b == nil {
		return shim.Error("queryMsgByID: query based on id failed!")
	}

	// 对查询到的状态进行反序列化
	var msg message_body.MsgBody
	err = json.Unmarshal(b, &msg)
	if err != nil {
		return shim.Error("queryMsgByID: unmarshal has error!")
	}

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(msg.MsgId)
	if err != nil {
		return shim.Error("queryMsgByID: failed to obtain historical data!")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []message_body.HistoryItem
	var hisMsg message_body.MsgBody
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("queryMsgByID: failed to obtain historical data!")
		}

		var historyItem message_body.HistoryItem
		historyItem.TxId = hisData.TxId
		err = json.Unmarshal(hisData.Value, &hisMsg)
		if err != nil {
			return shim.Error("queryMsgByID: failed to obtain historical data!")
		}

		if hisData.Value == nil {
			var empty message_body.MsgBody
			historyItem.MsgBody = empty
		} else {
			historyItem.MsgBody = hisMsg
		}

		historys = append(historys, historyItem)

	}

	msg.History = historys

	// 返回
	result, err := json.Marshal(msg)
	if err != nil {
		return shim.Error("queryMsgByID: unmarshal has error!")
	}
	return shim.Success(result)
}

// 根据身份证号更新信息
// args: educationObject
func (e *MsgChaincode) updateMsg(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var info message_body.MsgBody
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return shim.Error("updateMsg: unmarshal has error!")
	}

	// 根据身份证号码查询信息
	result, bl := GetMsg(stub, info.MsgId)
	if !bl {
		return shim.Error("updateMsg: query based on id failed!")
	}

	result.Sender = info.Sender
	result.Recipient = info.Recipient
	result.Body = info.Body
	result.DataHash = info.DataHash
	result.TimeDate = info.TimeDate

	_, bl = PutMsg(stub, result)
	if !bl {
		return shim.Error("updateMsg: an error occurred while saving the information!")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("updateMsg: success to update msg!"))
}

// 根据身份证号删除信息（暂不对外提供）
// args: entityID
func (e *MsgChaincode) delMsg(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("delMsg: unmarshal has error!")
	}

	/*var edu Education
	  result, bl := GetEduInfo(stub, info.EntityID)
	  err := json.Unmarshal(result, &edu)
	  if err != nil {
	      return shim.Error("反序列化信息时发生错误")
	  }*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("delMsg: an error occurred while deleting the information")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("delMsg: information deleted successfully"))
}
