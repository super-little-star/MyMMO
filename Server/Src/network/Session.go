package network

import ProtoMessage "mmo_server/ProtoMessage"

// TODO 用户数据实体

type GSession struct {
	response *ProtoMessage.NetMessage
}

func NewSession() *GSession {
	return &GSession{}
}

func (ns *GSession) Disconnected() {
	// TODO 会话断开后的处理
}

// GetNetResponse 获取Protobuf类型的Response
func (ns *GSession) GetNetResponse() *ProtoMessage.NetMessage {
	if ns.response == nil {
		ns.response = &ProtoMessage.NetMessage{}
	}
	return ns.response
}

// GetByteResponse 将Protobuf类型转换成Byte字节流
func (ns *GSession) GetByteResponse() []byte {
	if ns.response != nil {
		// TODO 消息处理
	}
	return nil
}
