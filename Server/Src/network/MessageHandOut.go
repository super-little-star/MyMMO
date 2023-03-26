package network

import (
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/singleton"
)

type GMessageHandOut struct {
}

func (mh *GMessageHandOut) HandOutResponse(sender *GConnection, response *ProtoMessage.NetMessageResponse) {
	if response == nil {
		return
	}
	if response.UserRegister != nil {
		singleton.Instance().MessageHandleCenter.TriggerEvents(sender, response.UserRegister)
	}
}

func (mh *GMessageHandOut) HandOutRequest(sender *GConnection, request *ProtoMessage.NetMessageRequest) {
	if request == nil {
		return
	}
	if request.UserRegister != nil {
		singleton.Instance().MessageHandleCenter.TriggerEvents(sender, request.UserRegister)
	}

}
