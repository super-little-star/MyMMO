package model

import (
	"errors"
	"mmo_server/ProtoMessage"
	"mmo_server/gameData"
	"mmo_server/network"
	"mmo_server/object"
)

var (
	ErrMapCharacterIsExist            = errors.New("map error:: character is exist")
	ErrMapConnectionCharacterNotFound = errors.New("map error:: connection character is not found")
)

type Map struct {
	data           *gameData.MapData
	charactersConn map[int32]*network.GConnection
}

func NewMap(data *gameData.MapData) *Map {
	return &Map{
		data:           data,
		charactersConn: make(map[int32]*network.GConnection),
	}
}

// Enter
//
//	@Description: 有新玩家进入
//	@receiver m
//	@param newOne 新玩家连接
//	@return error
func (m *Map) Enter(newOne *network.GConnection) error {
	response := newOne.Session().GetNetResponse()
	response.MapCharacterEnter = &ProtoMessage.MapCharacterEnterResponse{}

	if err := m.addCharacter(newOne); err != nil {
		response.MapCharacterEnter.Result = ProtoMessage.RESULT_FAILED
		response.MapCharacterEnter.Error = ProtoMessage.Error_MapCharacterIsExist
		newOne.SendResponse()
		return err
	}

	response.MapCharacterEnter.Result = ProtoMessage.RESULT_SUCCESS
	cs := response.MapCharacterEnter.Characters
	for _, v := range m.charactersConn {
		cs = append(cs, v.Session().CurCharacter.Proto)
		if v.Session().CurCharacter != newOne.Session().CurCharacter {
			// 通知地图其他玩家有新玩家进入
			if err := m.notifySomeCharacterEnter(v, newOne.Session().CurCharacter); err != nil {
				return nil
			}
		}
	}
	newOne.SendResponse()
	return nil
}

// addCharacter
//
//	@Description: 将玩家添加到地图map里
//	@receiver m
//	@param con
//	@return error
func (m *Map) addCharacter(con *network.GConnection) error {
	// 检查玩家是否已存在
	if con.Session().CurCharacter == nil {
		return ErrMapConnectionCharacterNotFound
	}

	if _, ok := m.charactersConn[con.Session().CurCharacter.Db.ID]; ok {
		return ErrMapCharacterIsExist
	}

	char := con.Session().CurCharacter

	char.MapId = m.data.ID
	m.charactersConn[char.Db.ID] = con

	return nil
}

// notifySomeCharacterEnter
//
//	@Description: 通知地图内玩家有角色进入
//	@receiver m
//	@param receiver 地图内玩家的链接
//	@param entrant 新进入的角色
//	@return error
func (m *Map) notifySomeCharacterEnter(receiver *network.GConnection, entrant *object.Character) error {
	if receiver.Session().GetNetResponse().MapCharacterEnter == nil {
		receiver.Session().GetNetResponse().MapCharacterEnter = &ProtoMessage.MapCharacterEnterResponse{}
	}
	response := receiver.Session().GetNetResponse().MapCharacterEnter

	response.MapId = m.data.ID
	response.Characters = append(response.Characters, entrant.Proto)

	receiver.SendResponse()
	return nil
}

func (m *Map) Update() {

}
