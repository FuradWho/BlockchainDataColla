package nats_service

import (
	"fmt"
	"github.com/FuradWho/BlockchainDataColla/fabricDeploy/common/micro_service"
	"github.com/asim/go-micro/v3/broker"
	log "github.com/sirupsen/logrus"
)

func SaveMsg(msg string) {

	option := micro_service.MicroOption.Option
	err := option.Broker.Publish("Msg", &broker.Message{
		Header: map[string]string{"type": msg},
		Body:   []byte("Msg broker nats"),
	})

	fmt.Println("222222")
	if err != nil {
		log.Errorln(err)
	}
}
