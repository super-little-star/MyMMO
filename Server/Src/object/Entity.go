package object

import (
	"mmo_server/ProtoMessage"
	"mmo_server/core"
)

type Entity struct {
	Id        int32
	Position  core.Vector3Int       // 位置
	Direction core.Vector3Int       // 朝向
	Proto     *ProtoMessage.PEntity // Entity 对应的Protobuf
}

func (e *Entity) SetPosition(pos core.Vector3Int, dir core.Vector3Int) {
	e.Position = pos
	e.Direction = pos
}

func (e *Entity) SetProto(proto *ProtoMessage.PEntity) {
	e.Proto = proto
}
