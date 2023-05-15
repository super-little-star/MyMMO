package network

import (
	"mmo_server/DB/Model"
	ProtoMessage "mmo_server/ProtoMessage"
)

// TODO 用户数据实体

type GSession struct {
	User   *Model.DbUser
	NetMsg *ProtoMessage.NetMessage
}

func NewSession() *GSession {
	return &GSession{}
}

func (ns *GSession) Disconnected() {
	if ns.User != nil {
		ConnectionManager().RemoveUser(ns.User.UID)
		ns.User = nil
	}
	ns.NetMsg = nil
}

// GetNetResponse 获取Protobuf类型的Response
func (ns *GSession) GetNetResponse() *ProtoMessage.NetMessageResponse {
	if ns.NetMsg == nil {
		ns.NetMsg = &ProtoMessage.NetMessage{
			Response: &ProtoMessage.NetMessageResponse{},
		}
	}
	return ns.NetMsg.Response
}

// GetByteResponse 将Protobuf类型转换成Byte字节流
func (ns *GSession) GetByteResponse() []byte {
	if ns.NetMsg != nil {
		// TODO 消息处理
	}
	return nil
}
