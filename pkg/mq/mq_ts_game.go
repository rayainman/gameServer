package mq

import (
	"encoding/json"
	"fmt"
	"gameserver/internal/game/model/player"
	"gameserver/internal/game/model/room"

	"github.com/streadway/amqp"
	"gitlab.diresoft.net/module/logger/logger"
	"gitlab.diresoft.net/server/lib.global/global/MQ"
)

type CreateRoomResult struct {
	Id string `json:"id"`
}

func (m_MQ *MQData) CreateRoom(mqDelivery *amqp.Delivery) MQ.ReturnType {

	id, ok := room.CreateRoom()

	if !ok {
		SendMQError(mqDelivery, fmt.Errorf("failed to create room"))
		return MQ.Cancel
	}

	createRoomResult := CreateRoomResult{
		Id: id,
	}

	err := json.Unmarshal(mqDelivery.Body, &createRoomResult)
	if err != nil {
		SendMQError(mqDelivery, fmt.Errorf("mqCreateRoom Get json error, %+v", err))
		return MQ.Cancel
	}

	logger.Info("%v", createRoomResult)

	return MQ.Ack //完成.
}

type EnterRoomRequest struct {
	RoomID    string     `json:"room_id"`
	Audiences []Audience `json:"audiences"`
}

type Audience struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
	Brand    string `json:"brand"`
	Site     string `json:"site"`
}

// type EnterRoomResult struct {
// 	Ok bool `json:"ok"`
// }

func (m_MQ *MQData) EnterRoom(mqDelivery *amqp.Delivery) MQ.ReturnType {

	enterRoomRequest := EnterRoomRequest{}

	err := json.Unmarshal(mqDelivery.Body, &enterRoomRequest)
	if err != nil {
		SendMQError(mqDelivery, fmt.Errorf("mqEnterRoom Get json error, %+v", err))
		return MQ.Cancel
	}

	logger.Info("%v", enterRoomRequest)

	roomID := enterRoomRequest.RoomID

	if room.GetRoom(roomID) == nil {
		SendMQError(mqDelivery, fmt.Errorf("room not found"))
		return MQ.Cancel
	}

	audiences := enterRoomRequest.Audiences
	for _, audience := range audiences {
		userName := audience.UserName
		userID := audience.UserID
		player := player.NewAudience(userID, userName, nil)

		if ok := room.GetRoom(roomID).AddAudience(player); !ok {
			return MQ.Cancel
		}

	}

	return MQ.Ack //完成.
}
