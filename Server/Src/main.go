package main

import (
	"fmt"
	"mmo_server/server"
	"mmo_server/utils/command"
	"mmo_server/utils/globalConfig"
	"mmo_server/utils/mlog"
	"mmo_server/utils/singleton"
)

func init() {
	globalConfig.Init()
	mlog.Init()
	singleton.InitSingleton()
}

func main() {
	fmt.Printf("Game Server Current Version [%s]\n", globalConfig.ProjectCfg.CurrVersion)
	gameServer := &server.GGameServer{}
	gameServer.Init()
	gameServer.Start()
	mlog.Info.Println("=====Game server is Running=====")
	command.Run()
	gameServer.Stop()
}
