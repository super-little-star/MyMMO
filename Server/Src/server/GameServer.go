package server

import (
	"mmo_server/network"
	"time"
)

type GGameServer struct {
	isRunning bool
	NetServer *network.GNetServer
}

// Init 初始化GameServer
func (gs *GGameServer) Init() {
	gs.isRunning = false
	gs.NetServer = &network.GNetServer{}
	gs.NetServer.Init("tcp", "127.0.0.1:7788")
}

// Start 开启GameServer逻辑
func (gs *GGameServer) Start() {
	gs.NetServer.Start()
	gs.isRunning = true
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
