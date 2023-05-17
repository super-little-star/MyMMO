package mConvert

import (
	"mmo_server/DB"
	"mmo_server/ProtoMessage"
	"mmo_server/manager"
	"mmo_server/network"
)

// Err2Protobuf
//
//	@Description: 将error装换成Protobuf的枚举
//	@param err
//	@return ProtoMessage.Error
func Err2Protobuf(err error) ProtoMessage.Error {
	switch err {
	case DB.ErrUserNameExist: // 用户名已存在
		return ProtoMessage.Error_RegisterUserNameExist

	case DB.ErrUserNotExist: // 用户不存在
		return ProtoMessage.Error_LoginUserNotExist
	case manager.ErrPasswordNotMatch: // 密码不正确
		return ProtoMessage.Error_LoginPasswordNotMatch
	case network.ErrUserIsOnline: // 用户已登录
		return ProtoMessage.Error_LoginUserIsOnline
	case DB.ErrCharacterNameExist: // 创建角色名已存在
		return ProtoMessage.Error_CreateCharacterNameExist
	default:
		return ProtoMessage.Error_None
	}
}
