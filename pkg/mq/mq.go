package mq

import (
	"fmt"
	"os"
	"sync"

	logger "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gitlab.diresoft.net/server/lib.global/global/MQ"
)

const (
	MatchServer     = "match"
	MatchGameServer = "tournament_game"
)

// enum Match Server 提供的功能.
const (
	mq_match_info  = "MatchInfo"
	mq_create_room = "CreateRoom"
)

// MQ.
type MQData struct {
	m_MQMutex sync.Mutex
	m_Receive *MQ.MQData //接收用.

	Connect      *MQ.MQData //跟 MQ 連接(發送用).
	ExchangeName string
	QueueName    string
	MyRoutingKey string
	Address      string
}

type MessageQueueNotify struct {
	CreateRoom func(*amqp.Delivery) MQ.ReturnType
	EnterRoom  func(*amqp.Delivery) MQ.ReturnType
	MatchInfo  func(*amqp.Delivery) MQ.ReturnType
}

var instance MQData = MQData{}

func GetInstance() *MQData {
	return &instance
}

func (m_MQ *MQData) SetAddress(address string) {
	m_MQ.Address = address
}

func (m_MQ *MQData) Init(MatchType string, notify *MessageQueueNotify) error {

	//MQ 接收用.
	m_MQ.m_Receive = MQ.Create(m_MQ.Address+"/", true)
	if m_MQ.m_Receive == nil {
		return fmt.Errorf("MQ create failed")
	}

	m_MQ.ExchangeName = "e"
	m_MQ.QueueName = "g" + MatchType
	m_MQ.MyRoutingKey = MatchType

	//註冊 Queue.
	if !m_MQ.m_Receive.Register(m_MQ.ExchangeName, m_MQ.QueueName, m_MQ.MyRoutingKey, 86400000) {
		return fmt.Errorf("MQ register failed")
	}

	//MQ 發送用.
	m_MQ.Connect = MQ.Create(m_MQ.Address+"/", true)
	if m_MQ.Connect == nil {
		return fmt.Errorf("connect to mq failed")
	}
	m_MQ.Connect.SetMessageAppId(m_MQ.QueueName) //夾帶以表示發送者身分.

	//接收 Queue 資料(前面都準備完之後才能開始接收，以免裡面要處理的東西需要上面功能).
	err := m_MQ.m_Receive.ReceiveManager(m_MQ.QueueName,
		MQ.ReceiveFunction{ //從 match 端來的訊息.
			mq_match_info:  notify.MatchInfo,
			mq_create_room: notify.CreateRoom,
		},
		mqOnReceive, mqOnClose)

	if err != nil {
		return err
	}

	return err
}

// 這裡配合 lib.global K8s 優雅退出，多一_InOutData參數，便於打印Logs
// 2022-03-18 來自igdMax的回覆
// 那個只是提供斷線當下的MQ資料
// 方便看log而已XD
// 我是會把他印出來看看啦
// 確認一下斷線時是在做什麼
func mqOnClose(_InQueueName string, _InOutData **amqp.Delivery) {
	//MQ 斷了就關掉程式.
	fmt.Printf("mqOnClose Queue:%s, %+v\n", _InQueueName, *_InOutData)
	os.Exit(1) //讓 K8S 重啟自己.
}

func mqOnReceive(_InData *amqp.Delivery) MQ.ReturnType {
	//收到不在定義內的其他資料.
	fmt.Println("msgser: Unknown MsgId -", _InData.MessageId)
	return MQ.Cancel
}

func (m_MQ *MQData) SendData(
	_InExchangeName string,
	_InRoutingKey string,
	_InMessageID string,
	_InType MQ.MessageType,
	_InCorrelationId string,
	_InHeaders amqp.Table,
	_InData []byte) {
	m_MQ.m_MQMutex.Lock()
	defer m_MQ.m_MQMutex.Unlock()

	if m_MQ.Connect.Send(_InExchangeName, _InRoutingKey, //m_MQ.ExchangeName
		_InMessageID,
		_InType,
		m_MQ.QueueName,
		_InCorrelationId,
		_InHeaders,
		_InData,
	) != nil {
		fmt.Println("MQ SendData error, reboot.")

		//MQ 斷了就關掉程式.
		os.Exit(1) //要重啟.
	}
}

func SendMQError(mqData *amqp.Delivery, err error) {
	logger.Error("%s", err.Error())

	//從 MQ 返回錯誤資訊.
	GetInstance().SendData("", mqData.ReplyTo,
		mqData.MessageId,
		MQ.Rep,
		mqData.CorrelationId,
		mqData.Headers,
		[]byte(err.Error()))
}
