package main

import (
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
	gameServer := &server.GGameServer{}
	gameServer.Init()
	gameServer.Start()
	mlog.Info.Println("=====Game server is Running=====")
	command.Run()
	gameServer.Stop()

}
