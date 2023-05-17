package object

import (
	"mmo_server/DB/DbObject"
	"mmo_server/ProtoMessage"
)

type Character struct {
	Creature Creature
	Db       *DbObject.DbCharacter
	Proto    *ProtoMessage.PCharacter // Character 对应的Protobuf
	MapId    int32
}

func NewCharacter(db *DbObject.DbCharacter) *Character {
	return &Character{
		Db: db,
	}
}
