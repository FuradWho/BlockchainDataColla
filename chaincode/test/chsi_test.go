package test

import (
	"encoding/json"
	"fmt"
	chaincode "github.com/FuradWho/BlockchainDataColla/chaincode/chaincode/eduChaincode"
	"github.com/FuradWho/BlockchainDataColla/chaincode/model/education"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"testing"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func TestInvoke(t *testing.T) {

	edu := &education.Education{
		Name:           "张三",
		Gender:         "男",
		Nation:         "汉",
		EntityID:       "101",
		Place:          "北京",
		BirthDay:       "1991年01月01日",
		EnrollDate:     "2009年9月",
		GraduationDate: "2013年7月",
		SchoolName:     "中国政法大学",
		Major:          "社会学",
		QuaType:        "普通",
		Length:         "四年",
		Mode:           "普通全日制",
		Level:          "本科",
		Graduation:     "毕业",
		CertNo:         "111",
		Photo:          "/static/phone/11.png",
	}

	b, err := json.Marshal(edu)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)

	cc := new(chaincode.EducationChaincode)
	stub := shimtest.NewMockStub("chsi", cc) // 创建MockStub对象
	// 调用Init接口，将a的值设为90
	// stub.MockInit("1", [][]byte{[]byte("a"), []byte("90")})
	// 调用get接口查询a的值
	res := stub.MockInvoke("1", [][]byte{[]byte("addEdu"), b, []byte("eventAddEdu")})
	fmt.Println("The value of a is", string(res.Payload))
	// 调用set接口设置a为100
	// stub.MockInvoke("1", [][]byte{[]byte("set"), []byte("a"), []byte("100")})
	// 再次查询a的值
	result, err := stub.GetQueryResult("{\"selector\":{\"docType\":\"eduObj\",\"CertNo\":\"111\",\"Name\":\"张三\"}}")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	res = stub.MockInvoke("1", [][]byte{[]byte("queryEduByCertNoAndName"), []byte("111"), []byte("张三")})
	fmt.Println("The new value of a is", string(res.Payload))

	// TODO
}
