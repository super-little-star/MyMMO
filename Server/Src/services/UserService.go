package services

import (
	"mmo_server/ProtoMessage"
	"mmo_server/manager"
	"mmo_server/network"
	"mmo_server/utils/err2protobuf"
	"mmo_server/utils/mlog"
)

type GUserService struct {
	manager manager.IUserManager
}

func (g *GUserService) Init() {
	network.LoginEvent[*ProtoMessage.RegisterRequest](g.OnUserRegister)
	network.LoginEvent[*ProtoMessage.LoginRequest](g.OnUserLogin)
	g.manager = manager.NewUserManager()
}
func (g *GUserService) Stop() {
	network.LogoffEvent[*ProtoMessage.RegisterRequest]()
	network.LogoffEvent[*ProtoMessage.LoginRequest]()
}

// OnUserRegister
//
//	@Description: 用户注册消息触发的事件
//	@receiver g
//	@param sender 发送者
//	@param msg 消息
func (g *GUserService) OnUserRegister(sender *network.GConnection, msg interface{}) {
	request, ok := msg.(*ProtoMessage.RegisterRequest)
	if !ok {
		mlog.Warning.Printf("Message[NUserRegisterRequest] 强转失败")
		return
	}

	mlog.Info.Printf("OnUserRegister:: UserName[%s] Password[%s]", request.UserName, request.Passward)
	// 写响应消息
	newMsg := &ProtoMessage.NetMessage{
		Response: &ProtoMessage.NetMessageResponse{
			Register: &ProtoMessage.RegisterResponse{},
		},
	}

	if err := g.manager.UserRegister(request.UserName, request.Passward); err != nil {
		newMsg.Response.Register.Result = ProtoMessage.RESULT_FAILED
		newMsg.Response.Register.Error = err2protobuf.Change(err)
		mlog.Error.Printf("User Service is error : %v", err)
	} else {
		newMsg.Response.Register.Result = ProtoMessage.RESULT_SUCCESS
		newMsg.Response.Register.Error = ProtoMessage.Error_None
	}

	sender.SendMsg(newMsg) // 将Response发送
}

// OnUserLogin
//
//	@Description: 用户登录触发事件
//	@receiver g
//	@param sender
//	@param msg
func (g *GUserService) OnUserLogin(sender *network.GConnection, msg interface{}) {
	request, ok := msg.(*ProtoMessage.RegisterRequest)
	if !ok {
		mlog.Warning.Printf("Message[NUserRegisterRequest] 强转失败")
		return
	}

	mlog.Info.Println("OnUserLogin:: UserName[%s] Password[%s]", request.UserName, request.Passward)

	newMsg := &ProtoMessage.NetMessage{
		Response: &ProtoMessage.NetMessageResponse{
			Login: &ProtoMessage.LoginResponse{},
		},
	}

	// 获取数据库数据
	dbUser, err := g.manager.UserLogin(request.UserName, request.Passward)
	if err != nil {
		newMsg.Response.Login.Result = ProtoMessage.RESULT_FAILED
		newMsg.Response.Login.Error = err2protobuf.Change(err)
		sender.SendMsg(newMsg)
		return
	}

	newMsg.Response.Login.User = &ProtoMessage.PUser{
		Uid:        dbUser.UID,
		Characters: []*ProtoMessage.PCharacter{},
	}

	for _, c := range dbUser.Characters {
		netCharacter := &ProtoMessage.PCharacter{
			Id:       c.ID,
			Name:     c.Name,
			Vocation: interface{}(c.Class).(ProtoMessage.VOCATION),
			Type:     ProtoMessage.CharacterType_Player,
		}
		newMsg.Response.Login.User.Characters = append(newMsg.Response.Login.User.Characters, netCharacter)
	}

	newMsg.Response.Login.Result = ProtoMessage.RESULT_SUCCESS
	newMsg.Response.Login.Error = ProtoMessage.Error_None

	sender.SendMsg(newMsg)
}
