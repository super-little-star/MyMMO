package services

import (
	"mmo_server/manager"
	"mmo_server/network"
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
)

type GUserService struct {
	manager manager.IUserManager
}

func (g *GUserService) Init() {
	network.LoginEvent[*ProtoMessage.NUserRegisterRequest](g.OnUserRegister)
	g.manager = manager.NewUserManager()
}
func (g *GUserService) Stop() {
	network.LogoffEvent[*ProtoMessage.NUserRegisterRequest]()
}

// OnUserRegister
//
//	@Description: 用户注册消息触发的事件
//	@receiver g
//	@param sender 发送者
//	@param msg 消息
func (g *GUserService) OnUserRegister(sender *network.GConnection, msg interface{}) {
	request, ok := msg.(*ProtoMessage.NUserRegisterRequest)
	if !ok {
		mlog.Warning.Printf("Message[NUserRegisterRequest] 强转失败")
		return
	}

	mlog.Info.Printf("OnUserRegister:: UserName[%s] Password[%s]", request.UserName, request.Passward)
	response := &ProtoMessage.NetMessage{}
	response.Response = &ProtoMessage.NetMessageResponse{}
	response.Response.UserRegister = &ProtoMessage.NUserRegisterResponse{}

	if err := g.manager.UserRegister(request.UserName, request.Passward); err != nil {
		response.Response.UserRegister.Result = ProtoMessage.RESULT_FAILED
		response.Response.UserRegister.Errormsg = err.Error()
	} else {
		response.Response.UserRegister.Result = ProtoMessage.RESULT_SUCCESS
		response.Response.UserRegister.Errormsg = ""
	}

	sender.SendMsg(response)
}
