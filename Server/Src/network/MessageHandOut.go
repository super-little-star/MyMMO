package network

import (
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
	"reflect"
)

type GMessageHandOut struct {
	messageEvents map[string]func(sender *GConnection, message interface{}) // 消息对应的世界处理
}

func (mh *GMessageHandOut) Init() {
	mh.messageEvents = make(map[string]func(sender *GConnection, message interface{}))
}

// HandOutResponse
//
//	@Description: 把Response发送给对应的事件进行处理
//	@receiver mh
//	@param sender
//	@param response
func (mh *GMessageHandOut) HandOutResponse(sender *GConnection, response *ProtoMessage.NetMessageResponse) {
	if response == nil {
		return
	}
	if response.UserRegister != nil {
		mh.TriggerEvents(sender, response)
	}
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
	if request.UserRegister != nil {
		mh.TriggerEvents(sender, request)
	}

}

// TriggerEvents
//
//	@Description: 根据所传入的消息触发对应的事件
//	@receiver m
//	@param sender 发送者
//	@param mes 消息
func (mh *GMessageHandOut) TriggerEvents(sender *GConnection, mes interface{}) {
	key := reflect.TypeOf(mes).Name()
	event, ok := mh.messageEvents[key]
	if ok {
		event(sender, mes)
	} else {
		mlog.Warning.Printf("message events the key[%s] is not find", key)
	}
}

// Login
//
//	@Description: 注册消息对应的事件
//	@receiver m
//	@param msg 消息
//	@param event 事件
func (mh *GMessageHandOut) Login(msg interface{}, event func(sender *GConnection, msg interface{})) {
	key := reflect.TypeOf(msg).Name()
	if mh.messageEvents[key] != nil {
		mh.messageEvents[key] = nil
	}
	mh.messageEvents[key] = event

}

// Logoff
//
//	@Description: 注销消息对应的事件
//	@receiver m
//	@param msg 消息
//	@param event 事件
func (mh *GMessageHandOut) Logoff(msg interface{}) {
	key := reflect.TypeOf(msg).Name()
	mh.messageEvents[key] = nil
}
