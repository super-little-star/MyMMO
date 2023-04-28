package err2protobuf

import (
	"mmo_server/DB"
	"mmo_server/ProtoMessage"
	"mmo_server/manager"
	"mmo_server/network"
)

// Convert
//
//	@Description: 将error装换成Protobuf的枚举
//	@param err
//	@return ProtoMessage.Error
func Convert(err error) ProtoMessage.Error {
	switch err {
	case DB.ErrUserNotExist: // 用户不存在
		return ProtoMessage.Error_LoginUserNotExist
	case DB.ErrUserNameExist: // 用户名已存在
		return ProtoMessage.Error_RegisterUserNameExist
	case manager.ErrPasswordNotMatch: // 密码不正确
		return ProtoMessage.Error_LoginPasswordNotMatch
	case network.ErrUserIsOnline:
		return ProtoMessage.Error_LoginUserIsOnline
	default:
		return ProtoMessage.Error_None
	}
}
