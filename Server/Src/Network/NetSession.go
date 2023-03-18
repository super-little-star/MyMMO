package Network

import ProtoMessage "mmo_server/protocol"

type GNetSession struct {
	response ProtoMessage.NetMessage
}

func (ns *GNetSession) GetResponse() []byte {
	return nil
}
