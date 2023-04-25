package network

import (
	"errors"
	"mmo_server/DB/Model"
	"mmo_server/utils/mlog"
)

var ErrConnIsExist = errors.New("connectionManager:: connection is exist")
var ErrUserIsOnline = errors.New("connectionManager:: user is online")

var connectionManager IConnectionManager

type IConnectionManager interface {
	Init()

	AddUser(uid int64, user *Model.DbUser) error
	RemoveUser(uid int64)

	GetConn(characterId int) *GConnection
	AddConn(characterId int, conn *GConnection) error
	RemoveConn(characterId int)
}

type GConnectionManager struct {
	connections map[int]*GConnection
	Users       map[int64]*Model.DbUser
}

func ConnectionManagerInit() {
	connectionManager = &GConnectionManager{}
	connectionManager.Init()
}
func ConnectionManager() IConnectionManager {
	return connectionManager
}

func (cm *GConnectionManager) Init() {
	cm.connections = make(map[int]*GConnection)
	cm.Users = make(map[int64]*Model.DbUser)
}

func (cm *GConnectionManager) AddUser(uid int64, user *Model.DbUser) error {
	if _, ok := cm.Users[uid]; ok {
		return ErrUserIsOnline
	} else {
		cm.Users[uid] = user
		return nil
	}
}
func (cm *GConnectionManager) RemoveUser(uid int64) {
	delete(cm.Users, uid)
}

func (cm *GConnectionManager) AddConn(characterId int, conn *GConnection) error {
	if _, ok := cm.connections[characterId]; !ok {
		cm.connections[characterId] = conn
		return nil
	} else {
		return ErrConnIsExist
	}
}

func (cm *GConnectionManager) RemoveConn(characterId int) {
	delete(cm.connections, characterId)
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
