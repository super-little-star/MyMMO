package model

import "mmo_server/core"

type Entity struct {
	Position  core.Vector3Int // 位置
	Direction core.Vector3Int // 朝向
	Proto     PEntity
}
