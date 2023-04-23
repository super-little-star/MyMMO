package network

import (
	ProtoMessage "mmo_server/ProtoMessage"
	"mmo_server/utils/mlog"
	"reflect"
)

var messageHandOut IMessageHandOut

type IMessageHandOut interface {
	Init()
	AddEvent(key string, event func(sender *GConnection, msg interface{}))
	RemoveEvent(key string)
	HandOutRequest(sender *GConnection, request *ProtoMessage.NetMessageRequest)
}

type GMessageHandOut struct {
	messageEvents map[string]func(sender *GConnection, message interface{}) // 消息对应的世界处理
}

func MessageHandOutInit() {
	messageHandOut = &GMessageHandOut{}
	messageHandOut.Init()
}

func MessageHandout() IMessageHandOut {
	return messageHandOut
}

func (mh *GMessageHandOut) Init() {
	mh.messageEvents = make(map[string]func(sender *GConnection, message interface{}))
}

// HandOutRequest
//
//	@Description: 把Request发给对应的事件处理
//	@receiver mh
//	@param sender
//	@param request
func (mh *GMessageHandOut) HandOutRequest(sender *GConnection, request *ProtoMessage.NetMessageRequest) {
	if request == nil {
		return
	}
	if request.Register != nil {
		mh.triggerEvents(sender, request.Register)
	}

}

func (mh *GMessageHandOut) AddEvent(key string, event func(sender *GConnection, msg interface{})) {
	_, ok := mh.messageEvents[key]
	if ok {
		mh.messageEvents[key] = nil
	}
	mh.messageEvents[key] = event
}

func (mh *GMessageHandOut) RemoveEvent(key string) {
	_, ok := mh.messageEvents[key]
	if ok {
		mh.messageEvents[key] = nil
	}
}

// triggerEvents
//
//	@Description: 根据所传入的消息触发对应的事件
//	@receiver m
//	@param sender 发送者
//	@param mes 消息
func (mh *GMessageHandOut) triggerEvents(sender *GConnection, mes interface{}) {
	key := reflect.TypeOf(mes).String()
	event, ok := mh.messageEvents[key]
	if ok {
		event(sender, mes)
		mlog.Info.Printf("Trigger Event[%s] success...", key)
	} else {
		mlog.Warning.Printf("message events the key[%s] is not find", key)
	}
}

// LoginEvent
//
//	@Description: 注册消息对应的事件
//	@param msg 消息
//	@param event 事件
func LoginEvent[T any](event func(sender *GConnection, msg interface{})) {
	var t T
	key := reflect.TypeOf(t).String()
	MessageHandout().AddEvent(key, event)
	mlog.Info.Printf("LoginEvent Message Event[%s]%v Success", key, event)
}

// LogoffEvent
//
//	@Description: 注销消息对应的事件
func LogoffEvent[T any]() {
	var t T
	key := reflect.TypeOf(t).String()
	MessageHandout().RemoveEvent(key)
	mlog.Info.Printf("LogoffEvent Message Event[%s] success", key)
}
