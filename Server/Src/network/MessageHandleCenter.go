package network

import (
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
	"sync/atomic"
	"time"
)

type IMessageHandleCenter interface {
	Init()
	Start(count int32)
	Stop()
	AcceptMessage(sender *GConnection, message *ProtoMessage.NetMessage)
}

type GMessagePackage struct {
	sender  *GConnection
	message *ProtoMessage.NetMessage
}

type GMessageHandleCenter struct {
	chanMessages         chan *GMessagePackage // 需要处理的信息Chan
	isRunning            bool                  // 消息处理中心是否在运行
	goroutinesCount      int32                 // 需要开启的Goroutines数量
	numRunningGoroutines int32                 // 正在运行的Goroutines数量
}

// Init 初始化消息处理中心
func (m *GMessageHandleCenter) Init() {
	m.chanMessages = make(chan *GMessagePackage, 1000)
	m.isRunning = false
	m.goroutinesCount = 0
	m.numRunningGoroutines = 0

}

// Start
//
//	@Description: 开启消息处理中心
//	@receiver m
//	@param count 需要开启Goroutine的数量
func (m *GMessageHandleCenter) Start(count int32) {
	m.isRunning = true

	// 把Goroutines数量控制在1000个以内
	if count <= 1 {
		m.goroutinesCount = 1
	} else if count > 1000 {
		m.goroutinesCount = 1000
	} else {
		m.goroutinesCount = count
	}

	for i := int32(1); i <= m.goroutinesCount; i++ {
		go m.messageDelivery()
	}

	for m.numRunningGoroutines < m.goroutinesCount {
		time.Sleep(100 * time.Millisecond)
	}
}

// Stop
//
//	@Description: 停止消息处理中心
//	@receiver m
func (m *GMessageHandleCenter) Stop() {
	close(m.chanMessages)
	m.isRunning = false

	for m.numRunningGoroutines > 0 {
		time.Sleep(100 * time.Millisecond)
	}
}

// messageDelivery 发送消息，将ChanMessage里的消息发送给对应的服务器进行处理
func (m *GMessageHandleCenter) messageDelivery() {
	//使用原子变量对运行的Goroutines计数
	atomic.AddInt32(&m.numRunningGoroutines, 1)
	mlog.Info.Printf("Message Delivery [No.%d]Goroutines is Start ...", m.numRunningGoroutines)
	for true {
		select {
		case pkg, ok := <-m.chanMessages:
			if !ok { // 如果chanMessages 通道被关闭，就是要停止分发信息
				atomic.AddInt32(&m.numRunningGoroutines, -1)
				mlog.Warning.Printf("chan messages is Closed and Message Delivery [No.%d]Goroutines is Stop !!!", m.numRunningGoroutines+1)
				return
			}
			// 把消息发送给Handout处理
			if pkg.message.Request != nil {
				Instance().MessageHandOut.HandOutRequest(pkg.sender, pkg.message.Request)
			}
			continue
		}
	}
}

// AcceptMessage
//
//	@Description: 接受消息，把消息放到消息处理中心的Chan里
//	@receiver m
//	@param sender 发送者
//	@param message 消息
func (m *GMessageHandleCenter) AcceptMessage(sender *GConnection, message *ProtoMessage.NetMessage) {
	ok := true
	if sender == nil {
		mlog.Warning.Printf("The message sender is null !!!")
		ok = false
	}
	if message == nil {
		mlog.Warning.Printf("The message is null !!!")
		ok = false
	}

	if !ok {
		return
	}

	m.chanMessages <- &GMessagePackage{
		sender:  sender,
		message: message,
	}
}
