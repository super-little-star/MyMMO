package main

import (
	"fmt"
	"mmo_server/server"
	"mmo_server/utils/command"
	"mmo_server/utils/globalConfig"
	"mmo_server/utils/mlog"
)

func init() {
	globalConfig.Init()
	mlog.Init()
}

func main() {
	fmt.Println(globalConfig.ProjectCfg.MaxPackageSize)
	gameServer := &server.GGameServer{IsRunning: false}
	gameServer.Init()
	gameServer.Start()
	mlog.Info.Println("=====Game server is Running=====")
	command.Run()
	gameServer.Stop()
}
