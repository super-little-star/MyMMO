package services

import (
	"mmo_server/network"
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
)

type GUserService struct {
}

func (g *GUserService) Init() {
	network.LoginEvent[*ProtoMessage.NUserRegisterRequest](g.OnUserRegister)
}
func (g *GUserService) Stop() {
	network.LogoffEvent[*ProtoMessage.NUserRegisterRequest]()
}

func (g *GUserService) OnUserRegister(sender *network.GConnection, msg interface{}) {
	request, ok := msg.(*ProtoMessage.NUserRegisterRequest)
	if !ok {
		mlog.Warning.Printf("Message[NUserRegisterRequest] 强转失败")
		return
	}
	mlog.Info.Printf("OnUserRegister:: UserName[%s] Password[%s]", request.UserName, request.Passward)
	response := &ProtoMessage.NetMessage{}
	response.Response = &ProtoMessage.NetMessageResponse{}
	response.Response.UserRegister = &ProtoMessage.NUserRegisterResponse{
		Result:   ProtoMessage.RESULT_SUCCESS,
		Errormsg: "",
	}

	sender.SendMsg(response)
}
