package main

import (
	"mmo_server/Server"
	"mmo_server/utils/command"
	"mmo_server/utils/globalConfig"
	"mmo_server/utils/mlog"
)

func init() {
	globalConfig.Init()
	mlog.Init()
}

func main() {
	server := &Server.GGameServer{IsRunning: false}
	server.Init()
	server.Start()
	mlog.Info.Println("=====Game server is Running=====")
	command.Run()
	server.Stop()
}
