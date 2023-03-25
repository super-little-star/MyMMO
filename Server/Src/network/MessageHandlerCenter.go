package network

import (
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
	"reflect"
	"sync/atomic"
	"time"
)

type MessagePackage struct {
	sender  *GConnection
	message *ProtoMessage.NetMessage
}

type MessageHandleCenter struct {
	ChanMessages           chan *MessagePackage
	messageEvents          map[string]func(sender *GConnection, message interface{})
	isRunning              bool
	goroutinesCount        uint32
	runningGoroutinesCount uint32
}

// Init 初始化消息处理中心
func (m *MessageHandleCenter) Init() {
	m.ChanMessages = make(chan *MessagePackage, 1000)
	m.messageEvents = make(map[string]func(sender *GConnection, message interface{}))
	m.isRunning = false
	m.goroutinesCount = 0
	m.runningGoroutinesCount = 0
}

// Start
//
//	@Description: 开启消息处理中心
//	@receiver m
//	@param count 需要开启Goroutine的数量
func (m *MessageHandleCenter) Start(count uint32) {
	m.isRunning = true

	if count <= 1 {
		m.goroutinesCount = 1
	} else if count > 1000 {
		m.goroutinesCount = 1000
	} else {
		m.goroutinesCount = count
	}

	for i := uint32(1); i <= m.goroutinesCount; i++ {
		go m.MessageDelivery()
	}

	for m.runningGoroutinesCount < m.goroutinesCount {
		time.Sleep(100 * time.Millisecond)
	}
}

// Stop
//
//	@Description: 停止消息处理中心
//	@receiver m
func (m *MessageHandleCenter) Stop() {
	close(m.ChanMessages)
	m.isRunning = false
	for m.runningGoroutinesCount > 0 {
		time.Sleep(100 * time.Millisecond)
	}
}

// MessageDelivery 发送消息，将ChanMessage里的消息发送给对应的服务器进行处理
func (m *MessageHandleCenter) MessageDelivery() {
	//使用原子变量对运行的Goroutines计数
	atomic.AddUint32(&m.runningGoroutinesCount, 1)
	defer atomic.AddUint32(&m.runningGoroutinesCount, -1)
	for {
		select {
		case _, ok := <-m.ChanMessages:
			if !ok {
				mlog.Warning.Printf("chan messages is Closed...")
				return
			}
			for m.isRunning {

			}
		}
	}
}

func (m *MessageHandleCenter) AcceptMessage(sender *GConnection, message *ProtoMessage.NetMessage) {
	m.ChanMessages <- &MessagePackage{
		sender:  sender,
		message: message,
	}
}

// TriggerEvents
//
//	@Description: 根据所传入的消息触发对应的事件
//	@receiver m
//	@param sender 发送者
//	@param mes 消息
func (m *MessageHandleCenter) TriggerEvents(sender *GConnection, mes interface{}) {
	key := reflect.TypeOf(mes).Name()
	event, ok := m.messageEvents[key]
	if ok {
		event(sender, mes)
	} else {
		mlog.Warning.Printf("message events the key[%s] is not find", key)
	}
}

// Register
//
//	@Description: 注册信息对应的事件
//	@receiver m
//	@param msg 信息
//	@param event 事件
func (m *MessageHandleCenter) Register(msg interface{}, event func(sender *GConnection, msg interface{})) {
	key := reflect.TypeOf(msg).Name()
	if m.messageEvents[key] != nil {
		m.messageEvents[key] = nil
	}
	m.messageEvents[key] = event

}

// Logoff
//
//	@Description: 注销信息对应的事件
//	@receiver m
//	@param msg 信息
//	@param event 事件
func (m *MessageHandleCenter) Logoff(msg interface{}) {
	key := reflect.TypeOf(msg).Name()
	m.messageEvents[key] = nil
}
