package mq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"gitlab.diresoft.net/module/logger/logger"
	"gitlab.diresoft.net/server/lib.global/global/MQ"
)

type NotifyResult struct {
	Platform string          `json:"Platform"`
	CameCode string          `json:"GameCode"`
	Currency string          `json:"Currency"`
	Account  []NotifyAccount `json:"Account"`
}

type NotifyAccount struct {
	Brand    string `json:"Brand"`
	Site     string `json:"Site"`
	PlayerID string `json:"PlayerID"`
}

func (m_MQ *MQData) MatchInfo(mqDelivery *amqp.Delivery) MQ.ReturnType {
	info := NotifyResult{}

	jsonerr := json.Unmarshal(mqDelivery.Body, &info)
	if jsonerr != nil {
		SendMQError(mqDelivery, fmt.Errorf("mqOnMatchInfo Get json error, %+v", jsonerr))
		return MQ.Cancel
	}

	logger.Info("%v", info)

	return MQ.Ack //完成.
}
