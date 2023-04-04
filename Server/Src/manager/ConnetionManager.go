package manager

import (
	"mmo_server/network"
	"mmo_server/utils/mlog"
)

type GConnectionManager struct {
	connections map[int]*network.GConnection
}

func (cm *GConnectionManager) Init() {
	cm.connections = make(map[int]*network.GConnection)
}

// GetConn
//
//	@Description: 获取连接
//	@receiver cm
//	@param characterId 角色ID
//	@return *network.GConnection 返回对应的链接
func (cm *GConnectionManager) GetConn(characterId int) *network.GConnection {
	c, ok := cm.connections[characterId]
	if ok {
		return c
	} else {
		mlog.Warning.Printf("Connection Manager character Id[%v] is not existent!!!", characterId)
		return nil
	}
}
