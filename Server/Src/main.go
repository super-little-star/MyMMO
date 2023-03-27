package main

import (
	"mmo_server/network"
	"mmo_server/server"
	"mmo_server/utils/command"
	"mmo_server/utils/globalConfig"
	"mmo_server/utils/mlog"
)

func init() {
	globalConfig.Init()
	mlog.Init()
	// 各个包单例模式初始化
	network.InitSingleton()
}

func main() {
	mlog.Info.Printf("Game Server Current Version [%s]\n", globalConfig.ProjectCfg.CurrVersion)
	gameServer := &server.GGameServer{}
	gameServer.Init()
	gameServer.Start()
	mlog.Info.Println("=====Game server is Running=====")
	command.Run()
	gameServer.Stop()
}
