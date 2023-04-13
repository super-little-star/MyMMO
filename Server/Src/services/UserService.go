package services

import (
	"mmo_server/DB"
	"mmo_server/ProtoMessage"
	"mmo_server/manager"
	"mmo_server/network"

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
	// 写响应消息
	newMsg := &ProtoMessage.NetMessage{
		Response: &ProtoMessage.NetMessageResponse{
			UserRegister: &ProtoMessage.NUserRegisterResponse{},
		},
	}

	if err := g.manager.UserRegister(request.UserName, request.Passward); err != nil {
		newMsg.Response.UserRegister.Result = ProtoMessage.RESULT_FAILED
		if err == DB.ErrUserNameExist { // 用户已存在
			newMsg.Response.UserRegister.Error = ProtoMessage.Error_UserNameExist
		}
		newMsg.Response.UserRegister.Error = ProtoMessage.Error_None
		mlog.Error.Printf("User Service is error : %v", err)
	} else {
		newMsg.Response.UserRegister.Result = ProtoMessage.RESULT_SUCCESS
		newMsg.Response.UserRegister.Error = ProtoMessage.Error_None
	}

	sender.SendMsg(newMsg) // 将Response发送
}
