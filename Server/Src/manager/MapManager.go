package manager

import (
	"errors"
	"mmo_server/gameData"
	"mmo_server/model"
)

var MapManager IMapManager

var ErrMapMgrMapIdNotFound = errors.New("map manager:: map id not found")

type IMapManager interface {
	Init()
	GetMap(mapId int32) (*model.Map, error)
	Update()
}

type GMapManager struct {
	maps map[int32]*model.Map
}

func InitMapManager() {
	MapManager = &GMapManager{}
	MapManager.Init()
}

func (mm *GMapManager) Init() {
	mm.maps = make(map[int32]*model.Map)

	// 新建所有地图对象
	for _, v := range gameData.DataManager.Maps {
		mm.maps[v.ID] = model.NewMap(v)
	}
}

// GetMap
//
//	@Description: 获取地图对象
//	@receiver mm
//	@param mapId 地图ID
//	@return *model.Map
//	@return error
func (mm *GMapManager) GetMap(mapId int32) (*model.Map, error) {
	res, ok := mm.maps[mapId]
	if !ok {
		return nil, ErrMapMgrMapIdNotFound
	}

	return res, nil
}

// Update
//
//	@Description: 更新所有地图的逻辑
//	@receiver mm
func (mm *GMapManager) Update() {
	for _, v := range mm.maps {
		v.Update()
	}
}
