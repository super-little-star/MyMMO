package main

import (
	"mmo_server/network"
	"mmo_server/services"
	"time"
)

type GGameServer struct {
	isRunning bool
	NetServer *network.GNetService
}

// Init 初始化GameServer
func (gs *GGameServer) Init() {
	gs.isRunning = false
	gs.NetServer = &network.GNetService{}
	gs.NetServer.Init("tcp", "127.0.0.1:7788")

	services.Instance().UserService.Init()
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
