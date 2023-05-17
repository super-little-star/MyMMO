package main

import (
	"mmo_server/DB"
	"mmo_server/gameData"
	"mmo_server/manager"
	"mmo_server/network"
	"mmo_server/services"
	"mmo_server/utils/globalConfig"
	"mmo_server/utils/mlog"
	"mmo_server/utils/xuid"
	"time"
)

type GGameServer struct {
	isRunning bool
	NetServer *network.GNetService
}

// Init 初始化GameServer
func (gs *GGameServer) Init() {
	gs.isRunning = false

	if err := gameData.Init(); err != nil {
		mlog.Error.Fatalf("Data Manager Init fail,error: %v !!!", err)
	}
	// 初始化网络服务
	gs.NetServer = &network.GNetService{}
	gs.NetServer.Init("tcp", "127.0.0.1:7788")

	// 初始化ID生成器
	if err := xuid.Init(10); err != nil {
		mlog.Error.Fatalf("ID Generator Init fail,error: %v !!!", err)
	}

	//初始化数据库
	if err := DB.Init(globalConfig.MySQLCfg.User, globalConfig.MySQLCfg.Password, globalConfig.MySQLCfg.IP, globalConfig.MySQLCfg.Port, globalConfig.MySQLCfg.DataBase); err != nil {
		mlog.Error.Fatalf("DB Init fail,error : %v !!!", err)
	} else {
		mlog.Info.Println("DB Init success ...")
	}

	manager.InitCharacterManager()

	services.UserService().Start()
}

// Start 开启GameServer逻辑
func (gs *GGameServer) Start() {
	gs.isRunning = true
	gs.NetServer.Start()

	go gs.update() // 开启一个携程 刷新游戏业务逻辑
}

// Stop 停止GamServer逻辑
func (gs *GGameServer) Stop() {
	gs.isRunning = false
	gs.NetServer.Stop()
}

// 更新GameServer逻辑处理
func (gs *GGameServer) update() {
	for gs.isRunning {
		time.Sleep(100 * time.Millisecond) // 每运行一次等待100毫秒
	}
}
