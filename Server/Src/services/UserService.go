package services

import (
	"mmo_server/ProtoMessage"
	"mmo_server/manager"
	"mmo_server/network"
	"mmo_server/utils/err2protobuf"
	"mmo_server/utils/mlog"
)

var userService IUserService

type IUserService interface {
	Start()
	Stop()
}

type GUserService struct {
	manager manager.IUserManager
}

func UserServiceInit() {
	userService = &GUserService{}
}

func UserService() IUserService {
	return userService
}

func (g *GUserService) Start() {
	g.manager = manager.NewUserManager()
	network.LoginEvent[*ProtoMessage.RegisterRequest](g.OnUserRegister)
	network.LoginEvent[*ProtoMessage.LoginRequest](g.OnUserLogin)
	network.LoginEvent[*ProtoMessage.CreateCharacterRequest](g.OnCreateCharacter)
}
func (g *GUserService) Stop() {
	network.LogoffEvent[*ProtoMessage.RegisterRequest]()
	network.LogoffEvent[*ProtoMessage.LoginRequest]()
	network.LogoffEvent[*ProtoMessage.CreateCharacterRequest]()
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
		newMsg.Response.Register.Error = err2protobuf.Convert(err)
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
	request, ok := msg.(*ProtoMessage.LoginRequest)
	if !ok {
		mlog.Warning.Printf("Message[NUserRegisterRequest] 强转失败")
		return
	}

	mlog.Info.Printf("OnUserLogin:: UserName[%s] Password[%s]", request.UserName, request.Passward)

	newMsg := &ProtoMessage.NetMessage{
		Response: &ProtoMessage.NetMessageResponse{
			Login: &ProtoMessage.LoginResponse{},
		},
	}

	// 获取数据库数据
	dbUser, err := g.manager.UserLogin(request.UserName, request.Passward)
	if err != nil {
		newMsg.Response.Login.Result = ProtoMessage.RESULT_FAILED
		newMsg.Response.Login.Error = err2protobuf.Convert(err)
		sender.SendMsg(newMsg)
		return
	}

	// 是否已经存在该User的链接
	if err := network.ConnectionManager().AddUser(dbUser); err != nil {
		newMsg.Response.Login.Result = ProtoMessage.RESULT_FAILED
		newMsg.Response.Login.Error = err2protobuf.Convert(err)
		sender.SendMsg(newMsg)
		return
	}
	sender.Session().User = dbUser

	newMsg.Response.Login.User = &ProtoMessage.PUser{
		Uid:        dbUser.UID,
		Characters: []*ProtoMessage.PCharacter{},
	}

	// 把dbUser的角色信息转换成protobuf
	for _, c := range dbUser.Characters {
		netCharacter := &ProtoMessage.PCharacter{
			Id:    c.ID,
			Name:  c.Name,
			Class: ProtoMessage.CharacterClass(c.Class),
			Level: c.Level,
		}
		newMsg.Response.Login.User.Characters = append(newMsg.Response.Login.User.Characters, netCharacter)
	}

	newMsg.Response.Login.Result = ProtoMessage.RESULT_SUCCESS
	newMsg.Response.Login.Error = ProtoMessage.Error_None

	sender.SendMsg(newMsg)
}

func (g *GUserService) OnCreateCharacter(sender *network.GConnection, msg interface{}) {
	request, ok := msg.(*ProtoMessage.CreateCharacterRequest)
	if !ok {
		mlog.Warning.Println("Message[CreateCharacterRequest]强转失败")
		return
	}

	mlog.Info.Printf("OnCreateCharacter:: Name[%v] Class[%v]", request.Name, request.CharacterClass)

	newCharacters, err := g.manager.CreateCharacter(
		sender.Session().User.UID,
		request.Name,
		int(request.CharacterClass),
	)

	sender.Session().User.Characters = newCharacters

	sender.Session().GetNetResponse().CreateCharacter = &ProtoMessage.CreateCharacterResponse{}

	if err != nil {
		sender.Session().GetNetResponse().CreateCharacter.Result = ProtoMessage.RESULT_FAILED
		sender.Session().GetNetResponse().CreateCharacter.Error = err2protobuf.Convert(err)
		sender.SendResponse()
		return
	}

	sender.Session().GetNetResponse().CreateCharacter.Result = ProtoMessage.RESULT_SUCCESS
	var netCharacters []*ProtoMessage.PCharacter
	for _, c := range newCharacters {
		netChar := &ProtoMessage.PCharacter{
			Id:    c.ID,
			Name:  c.Name,
			Class: ProtoMessage.CharacterClass(c.Class),
			Level: c.Level,
		}
		netCharacters = append(netCharacters, netChar)
	}

	sender.Session().GetNetResponse().CreateCharacter.Characters = netCharacters
	sender.SendResponse()

}
