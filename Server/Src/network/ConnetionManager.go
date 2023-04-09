package network

import (
	"mmo_server/utils/mlog"
)

type IConnectionManager interface {
	Init()
	GetConn(characterId int) *GConnection
}

type GConnectionManager struct {
	connections map[int]*GConnection
}

func (cm *GConnectionManager) Init() {
	cm.connections = make(map[int]*GConnection)
}

// GetConn
//
//	@Description: 获取连接
//	@receiver cm
//	@param characterId 角色ID
//	@return *network.GConnection 返回对应的链接
func (cm *GConnectionManager) GetConn(characterId int) *GConnection {
	c, ok := cm.connections[characterId]
	if ok {
		return c
	} else {
		mlog.Warning.Printf("Connection Manager character Id[%v] is not existent!!!", characterId)
		return nil
	}
}
