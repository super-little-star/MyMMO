package services

import (
	"mmo_server/network"
	ProtoMessage "mmo_server/protocol"
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
	response := &ProtoMessage.NetMessage{}
	response.Response = &ProtoMessage.NetMessageResponse{}
	response.Response.UserRegister = &ProtoMessage.NUserRegisterResponse{
		Result:   ProtoMessage.NRESULT_SUCCESS,
		Errormsg: "",
	}

	sender.SendMsg(response)
}
